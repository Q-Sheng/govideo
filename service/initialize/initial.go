package initialize

import (
	"service/global"
	"service/initialize/config"
	mygorm "service/initialize/gorm"
	"service/initialize/router"
)

//// ConfigInfo 配置信息，与配置文件config.ini文件相关
//var configInfo = new(global.ConfigInformation)
//
//// Db 根据配置信息初始化Db
//var Db = new(gorm.DB)
//
//// Router
//var Router = new(gin.Engine)

func InitServe() {
	// 初始化配置信息
	config.InitConfig(global.ConfigInfo)

	// 初始化数据库,初始化配置信息需要在前
	global.Db = mygorm.InitGorm(global.ConfigInfo)

	// 初始化Router，并返回做作为全局，并后续启动
	global.Router = router.InitRouter()
}
