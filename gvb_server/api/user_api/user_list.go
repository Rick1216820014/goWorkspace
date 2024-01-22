package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desensitization"
	"gvb_server/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {

	//如何判断是管理员 用中间件实现
	//token := c.Request.Header.Get("token")
	//if token == "" {
	//	res.FailWithMessage("未携带token", c)
	//}
	//claims, err := jwts.ParseToken(token)
	//if err != nil {
	//	res.FailWithMessage("token错误", c)
	//	return
	//}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var users []models.UserModel

	list, count, _ := common.Comlist(models.UserModel{}, common.Option{
		PageInfo: page,
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		//脱敏
		user.Tel = desensitization.DesensitizationTel(user.Tel)
		user.Email = desensitization.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
