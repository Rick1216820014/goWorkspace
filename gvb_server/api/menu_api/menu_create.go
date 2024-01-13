package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"` // 标题
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`   // 路径或别名
	Slogan        string      `json:"slogan"`                                                 // slogan
	Abstract      ctype.Array `json:"abstract"`                                               // 简介
	AbstractTime  int         `json:"abstract_time"`                                          // 简介的切换时间
	BannerTime    int         `json:"banner_time"`                                            // 菜单图片的切换时间 为 0 表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`   // 菜单的顺序
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                            //具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("菜单标题或路径重复", c)
		return

	}
	//创建bannner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
	}
	//给第三张表入库
	if len(cr.ImageSortList) == 0 {
		res.FailWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel

	for _, sort := range cr.ImageSortList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)

}
