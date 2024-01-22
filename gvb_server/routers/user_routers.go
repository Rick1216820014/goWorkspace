package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) UserRouter() {

	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateView)
	router.POST("update_password", middleware.JwtAuth(), app.UserUpdatePasswordView)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
}
