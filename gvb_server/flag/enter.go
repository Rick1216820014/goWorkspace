package flag

import (
	sys_flag "flag"
	"fmt"
)

type Oiption struct {
	//Version bool
	DB   bool
	User string //-u admin -u user
}

// 解析命令行参数
func Parse() Oiption {
	//version := sys_flag.Bool("v", false, "项目版本")
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")

	sys_flag.Parse()

	return Oiption{

		DB:   *db,
		User: *user,
	}
}

// 是否停止web项目
func IsWebStop(option Oiption) bool {
	fmt.Println("命令行调用-db", option.DB)
	if option.DB {
		//返回true表示web已经停止可以迁移
		return true
	}
	if option.User == "admin" || option.User == "user" {
		return true
	}
	return false
}

func SwitchOption(option Oiption) {
	if option.DB {
		Makemigrations()
		//fmt.Println("调试")
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}
	//sys_flag.Usage()
}
