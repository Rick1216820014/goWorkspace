package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	//图片上传白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

// 上传多个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	//判断路径是否存在
	//不存在就创建
	basepath := global.Config.Upload.Path
	_, err = os.ReadDir(basepath)
	if err != nil {
		err = os.MkdirAll(basepath, fs.ModePerm)
		global.Log.Error(err)
		global.Log.Println("本地不存在该路径，已创建")
	}

	var resList []FileUploadResponse

	for _, file := range fileList {
		fileName := file.Filename
		//filePath := path.Join("uploads", file.Filename)
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])

		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       "不支持上传该类型文件",
			})
			continue
		}

		filePath := path.Join(basepath, fileName)

		//判断大小

		size := float64(file.Size) / float64(1024*1024)

		if size >= float64(global.Config.Upload.Size) {

			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超出设定大小，当前大小为：%.2fMB，设定大小为:%dMB ", size, global.Config.Upload.Size),
			})
			continue
		}
		//读取图转为字节切片列表，转为md5值存入数据库，作为哈希值，用于判断文件是否重复上传
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)

		imageHash := utils.Md5(byteData)

		//去数据库中判断是否存在对应的哈希值
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash=?", imageHash).Error
		if err == nil {
			//存在图片对应的哈希值
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "图片已上传",
			})
			continue
		}

		//上传到七牛
		if global.Config.QiNiu.Enable {
			filePath, err = qiniu.UploadImage(byteData, fileName, "gvb")
			if err != nil {
				global.Log.Error(err)
				continue
			}
			resList = append(resList, FileUploadResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "上传七牛成功",
			})
			global.DB.Create(&models.BannerModel{

				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.QiNiu,
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
		//图片入库
		global.DB.Create(&models.BannerModel{

			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: ctype.Local,
		})
		//fmt.Println(filePath, float64(file.Size)/float64(1024*1024))

	}

	//单个上传
	//fileheader, err := c.FormFile("image")
	res.OkWithData(resList, c)
}
