package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwts"
)

func (UserApi) LogoutView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")

	//fmt.Println(claims.ExpiresAt)

	//需要计算过期时间距离现在的过期时间
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
