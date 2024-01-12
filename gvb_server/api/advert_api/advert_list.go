package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"
)

func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	//err := c.ShouldBindJSON(&cr)
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		global.Log.Error(err)
		return
	}
	referer := c.GetHeader("Referer")
	isShow := true
	//admin可以查看所有，隐藏不显示的也需要能看到
	if strings.Contains(referer, "admin") {
		isShow = false
	}

	list, count, _ := common.Comlist(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
