package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

// 图片列表
type Option struct {
	models.PageInfo
	Debug bool
}

func Comlist[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		//展示部分日志
		//logger.Interface 是 GORM 框架定义的一个接口类型，用于抽象不同的日志记录器实现。
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	//分页，limit每页多少行数据，offset从第几页开始(数据偏移量)
	//例如，假设你有一个总共有100条数据的列表，每页显示10条数据。
	//那么，如果你想获取第2页的数据，你需要设置 offset 为 (页码 - 1) * 每页数量，也就是 (2 - 1) * 10 = 10。
	//这样，查询结果将从第11条数据开始获取（也就是page=2,从第二页开始展示），获取10条数据，即第11到第20条数据。
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
