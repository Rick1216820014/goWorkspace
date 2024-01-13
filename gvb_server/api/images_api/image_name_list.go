package images_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`                // 图片路径
	Name string `gorm:"size:38" json:"name"` // 图片名称
}

// 主要用于菜单
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	DB := global.DB
	DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)

}
