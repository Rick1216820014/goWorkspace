package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	//fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘻嘻嘻")
	global.Log.Error("嘻嘻嘻")
	global.Log.Info("嘻嘻嘻")

	//连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
}
