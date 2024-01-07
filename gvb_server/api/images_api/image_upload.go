package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"
)

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

	var resList []image_ser.FileUploadResponse
	for _, file := range fileList {

		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		//成功
		if !global.Config.QiNiu.Enable {
			//本地保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.FileName = file.Filename
				serviceRes.IsSuccess = false

				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
		//fmt.Println(filePath, float64(file.Size)/float64(1024*1024))

	}

	//单个上传
	//fileheader, err := c.FormFile("image")
	res.OkWithData(resList, c)
}
