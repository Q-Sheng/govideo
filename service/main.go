package main

import (
	"fmt"
	"service/global"
	"service/initialize"
)

func main() {

	// todo 测试数据库，编写路由，初步打通数据与前端，测试是否能全局使用Db
	// 初始化
	initialize.InitServe()
	//初始化数据库,
	//检查global的Db是否已经初始化
	err := global.Db.Error
	if err != nil {
		fmt.Printf("err:", err)
	}
	//fmt.Printf("err:", err)
	//fmt.Println(global.Db.Statement)
	// 路由
	//router.InitRouter()
	//
	global.Router.Run(global.ConfigInfo.Serconfig.HttpPort)
}
