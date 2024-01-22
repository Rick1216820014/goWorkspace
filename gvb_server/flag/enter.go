package flag

import (
	sys_flag "flag"
	"fmt"
	"github.com/fatih/structs"
)

type Oiption struct {
	//Version bool
	DB   bool
	User string //-u admin -u user
	ES   string //-es create  -es delete
}

// 解析命令行参数
func Parse() Oiption {
	//version := sys_flag.Bool("v", false, "项目版本")
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "es操作")

	sys_flag.Parse()

	return Oiption{

		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// 是否停止web项目
func IsWebStop(option Oiption) (f bool) {
	fmt.Println("命令行调用-db", option.DB)
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

func SwitchOption(option Oiption) {
	if option.DB {
		Makemigrations()
		//fmt.Println("调试")
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}

	if option.ES == "create" {
		fmt.Println("es调试")
		EsCreateIndex()
	}
	//sys_flag.Usage()
}
