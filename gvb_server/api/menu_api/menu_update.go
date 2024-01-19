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
	str_id := c.Param("id")
	num, _ := strconv.ParseUint(str_id, 10, 64)

	// 将 uint64 类型转换为 uint 类型
	id := uint(num)
	var menuModel models.MenuModel
	var menuBannerModel models.MenuBannerModel
	bannerMenuList := []models.MenuBannerModel{}

	count := global.DB.Where("id = ?", id).Find(&menuModel).RowsAffected
	if count == 0 {
		res.FailWithMessage("不存在的菜单id", c)
		return
	}
	if len(cr.ImageSortList) > 0 {
		//mbs=menuid+bannerid+sort
		//把待更新的图片列表保存为切片
		mbsList := []models.MenuBannerModel{}
		for _, v := range cr.ImageSortList {
			mbsList = append(mbsList, models.MenuBannerModel{
				MenuID:   id,
				BannerID: v.ImageID,
				Sort:     v.Sort,
			})
		}

		count := global.DB.Where("menu_id = ?", id).Find(&menuBannerModel).Scan(&bannerMenuList).RowsAffected
		if count == 0 {
			global.Log.Println("menuBannerModel对应id下没有图片")
			//那就不用对比了，直接全更新
			err = global.DB.Debug().Create(&mbsList).Error
			if err != nil {
				res.FailWithMessage("图片存储异常", c)
				global.Log.Error(err)
				return
			}
		} else {
			create_mbsList := []models.MenuBannerModel{}
			update_mbsList := []models.MenuBannerModel{}
			delete_mbsList := []models.MenuBannerModel{}
			for i := 0; i < len(mbsList); i++ {

				for j := 0; j < len(bannerMenuList); j++ {
					//bannerid相同但是sort不同，记录在更新切片中
					if mbsList[i].BannerID == bannerMenuList[j].BannerID &&
						mbsList[i].Sort != bannerMenuList[j].Sort {
						update_mbsList = append(update_mbsList, models.MenuBannerModel{
							MenuID:   id,
							BannerID: mbsList[i].BannerID,
							Sort:     mbsList[i].Sort,
						})
						break
					} else if mbsList[i].BannerID == bannerMenuList[j].BannerID &&
						mbsList[i].Sort == bannerMenuList[j].Sort {
						//防止出现最后一次循环找到相同的切片，并且保存到插入切片中
						break
					}
					//如果确定数据库中没有这个图片，记录在插入切片中
					if j == len(bannerMenuList)-1 {
						create_mbsList = append(create_mbsList, models.MenuBannerModel{
							MenuID:   id,
							BannerID: mbsList[i].BannerID,
							Sort:     mbsList[i].Sort,
						})
					}
				}

			}
			//反向查找，数据库中的图片是否包含在新的图片列表中，如果没有就加入删除切片中
			for i := 0; i < len(bannerMenuList); i++ {

				for j := 0; j < len(mbsList); j++ {
					if mbsList[j].BannerID == bannerMenuList[i].BannerID &&
						mbsList[j].Sort == bannerMenuList[i].Sort {
						break
					} else if mbsList[j].BannerID == bannerMenuList[i].BannerID &&
						mbsList[j].Sort != bannerMenuList[i].Sort {
						break
					}
					if j == len(mbsList)-1 {
						delete_mbsList = append(delete_mbsList, models.MenuBannerModel{
							MenuID:   id,
							BannerID: bannerMenuList[i].BannerID,
							Sort:     bannerMenuList[i].Sort,
						})
					}
				}
			}
			if len(create_mbsList) != 0 {
				global.DB.Debug().Create(&create_mbsList)
			}

			if len(update_mbsList) != 0 {
				for _, data := range update_mbsList {
					global.DB.Debug().Model(&models.MenuBannerModel{}).Where("menu_id = ? and banner_id = ?", data.MenuID, data.BannerID).Updates(data)
				}

			}
			if len(delete_mbsList) != 0 {
				for _, data := range delete_mbsList {
					global.DB.Debug().Where("menu_id = ? and banner_id = ?", data.MenuID, data.BannerID).Delete(&models.MenuBannerModel{}, data)
				}
			}
		}
	} else {
		//如果更新的是空图片列表，就是删除关联的所有图片
		err = global.DB.Take(&menuModel, id).Error
		if err != nil {
			res.FailWithMessage("不存在该菜单", c)
		}
		global.DB.Model(&menuModel).Association("Banners").Clear()
	}
	// 普通更新
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Debug().Model(&menuModel).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
