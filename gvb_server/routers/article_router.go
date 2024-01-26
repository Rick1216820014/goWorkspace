package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("articles/:id", app.ArticleDetailView)
	router.GET("articles/calendar", app.ArticleCalenderView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles", middleware.JwtAdmin(), app.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAdmin(), app.ArticleRemoveView)
}
