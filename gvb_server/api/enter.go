package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/user_api"
	"gvb_server/test"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	TestApi     test.TestApi
	UserApi     user_api.UserApi
	ArticleApi  article_api.ArticleApi
}

// 全局变量ApiGroupApp,把ApiGroup实例化
var ApiGroupApp = new(ApiGroup)
