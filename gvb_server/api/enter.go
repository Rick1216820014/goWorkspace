package api

import (
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

// 全局变量ApiGroupApp,把ApiGroup实例化
var ApiGroupApp = new(ApiGroup)
