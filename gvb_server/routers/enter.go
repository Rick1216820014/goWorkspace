package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	//*gin.Engine
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//路由分层
	//系统配置api
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()

	return router

}

//router.GET("", func(c *gin.Context) {
//	c.String(200, "我爱小宝")
//
//})
