package menu_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"strconv"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	fmt.Println(id)
	var menuModel models.MenuModel
	var menuBannerModel models.MenuBannerModel
	var bannerMenuList []models.MenuBannerModel
	//bannerMenuList对应菜单id下的所有图片
	global.DB.Find(&menuBannerModel).Where("menu_id = ?", id).Scan(&bannerMenuList)
	//var bannerOld []models.MenuBannerModel
	bannnerMap := make(map[string]uint)
	//首先确保图片列表不为空
	if len(cr.ImageSortList) > 0 {
		//遍历输入的所有待更新的图片
		for _, v := range cr.ImageSortList {
			str := strconv.Itoa(int(v.ImageID))
			bannnerMap[str] = uint(v.Sort)
		}
		fmt.Println(bannnerMap)
		//遍历数据库中的菜单图片，查找是否有这张图片
		for _, banner := range bannerMenuList {
			strID := strconv.Itoa(int(banner.BannerID))
			_, exist := bannnerMap[strID]
			//若存在，直接更新编号
			if exist {
				global.DB.Debug().Model(&menuBannerModel).Where("menu_id = ? and banner_id = ?", id, banner.BannerID).Update("sort", bannnerMap[strID])
				continue
			} else {
				// 操作第三张表
				var bannerList []models.MenuBannerModel
				for _, sort := range cr.ImageSortList {
					bannerList = append(bannerList, models.MenuBannerModel{
						MenuID:   menuModel.ID,
						BannerID: sort.ImageID,
						Sort:     sort.Sort,
					})
				}
				err = global.DB.Create(&bannerList).Error
				if err != nil {
					global.Log.Error(err)
					res.FailWithMessage("创建菜单图片失败", c)
					return
				}
			}

		}
	}
	// 普通更新
	maps := structs.Map(&cr)
	fmt.Println("id", id)
	err = global.DB.Debug().Model(&menuModel).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)

}
