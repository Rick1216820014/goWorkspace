package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("该广告不存在", c)
		return
	}

	err = global.DB.Model(&advert).Updates(map[string]any{
		"Title":  cr.Title,
		"Href":   cr.Href,
		"Images": cr.Images,
		"IsShow": cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OkWithMessage("广告成功", c)
}
