package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MenuNameResponse struct {
	ID    uint
	Title string
	Path  string
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
