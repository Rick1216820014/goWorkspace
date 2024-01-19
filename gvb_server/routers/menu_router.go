package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {

	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.PUT("menus/:id", app.MenuUpdateView)
	router.DELETE("menus", app.MenuRemoveView)
	//router.PUT("testmenus/:id", api.ApiGroupApp.TestApi.MenuUpdateView)

}
