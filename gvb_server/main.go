package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//连接redis
	global.Redis = core.ConnectRedis()
	//连接es
	global.ESClient = core.EsConnect()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatal(err.Error())
	}

}
