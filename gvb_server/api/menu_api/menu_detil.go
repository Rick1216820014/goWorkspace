package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuDetilView(c *gin.Context) {
	//先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Debug().Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	//查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Debug().Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}

	res.OkWithData(menuResponse, c)
}
