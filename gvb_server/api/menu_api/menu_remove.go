package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		global.Log.Error(err)
		return
	}

	var menuList []models.MenuModel
	//Find()第一个参数是指针类型，需要传入一个切片的指针，否则会报错：panic recovered:reflect: reflect.Value.Set using unaddressable value
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
	}

	res.OkWithMessage(fmt.Sprintf("删除%d个菜单", count), c)
}
