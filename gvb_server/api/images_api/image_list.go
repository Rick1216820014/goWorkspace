package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {

	var cr models.PageInfo
	//通过 ShouldBindQuery 函数将前端传递的查询参数绑定到 cr 变量上
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.Comlist(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	//res.OkWithData(gin.H{"count": count, "list": imageList}, c)
	res.OkWithList(list, count, c)
	return

}
