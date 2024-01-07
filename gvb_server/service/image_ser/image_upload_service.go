package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
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

// ImageUploadService文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename

	res.FileName = fileName

	//文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])

	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	basepath := global.Config.Upload.Path

	//判断文件大小
	filePath := path.Join(basepath, fileName)

	size := float64(file.Size) / float64(1024*1024)

	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超出设定大小，当前大小为：%.2fMB，设定大小为:%dMB ", size, global.Config.Upload.Size)

		return
	}

	//读取文件内容 hash
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
		//找到了
		//存在图片对应的哈希值
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	//上传到七牛

	res.Msg = "图片上传成功"
	res.IsSuccess = true
	filePath = "/" + filePath
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu

	}
	//图片入库
	err = global.DB.Create(&models.BannerModel{

		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.Msg = "图片入库出错:" + err.Error()
		res.FileName = fileName
		return
	}
	return
}
