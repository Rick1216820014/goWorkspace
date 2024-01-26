package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
)

// json:"id"：表示在 JSON 序列化过程中，该字段的名称应为 id。
// form:"id"：表示在表单数据绑定过程中，该字段的名称应为 id。
// uri:"id"：表示在 URI 参数绑定过程中，该字段的名称应为 id。
type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
