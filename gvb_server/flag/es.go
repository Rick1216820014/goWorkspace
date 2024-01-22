package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func EsCreateIndex() {
	err := models.ArticleModel{}.CreateIndex()
	if err != nil {
		global.Log.Error(err)
		global.Log.Warn("es索引生成失败")
		return
	}
	global.Log.Info("es生成索引成功")

}
