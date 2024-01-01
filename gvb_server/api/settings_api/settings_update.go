package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	global.Log.Println("112")
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)

	if err != nil {
		global.Log.Println("114")
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":

		var info config.SiteInfo
		global.Log.Println(info)
		err = c.ShouldBindJSON(&info)
		if err != nil {
			global.Log.Println(err)
			global.Log.Println("116")
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Log.Println("111")
		global.Config.SiteInfo = info

	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwy
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwy = info
	default:

		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
	res.Okwith(c)
}
