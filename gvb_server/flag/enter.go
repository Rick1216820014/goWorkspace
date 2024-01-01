package flag

import sys_flag "flag"

type Oiption struct {
	//Version bool
	DB bool
}

// 解析命令行参数
func Parse() Oiption {
	//version := sys_flag.Bool("v", false, "项目版本")
	db := sys_flag.Bool("db", false, "初始化数据库")

	sys_flag.Parse()
	return Oiption{

		DB: *db,
	}
}

// 是否停止微web项目
func IsWebStop(option Oiption) bool {
	if option.DB {
		//返回true表示web已经停止可以迁移
		return true
	}
	return false
}

func SwitchOption(option Oiption) {
	if option.DB {
		Makemigrations()
	}

}
