package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"service/global"
)

/*
	配置文件config.ini的全局读取全局变量
*/

// var ConfigInfo = new(ConfigInformation)
var err error
var inifile *ini.File

func InitConfig(ConfigInfo *global.ConfigInformation) {
	//读取ini文件
	inifile, err = ini.Load("./config/config.ini")
	if err != nil {
		log.Printf("无法读取配置文件config.ini:", err)
	}

	err = inifile.Section("database").MapTo(&ConfigInfo.Dbconfig)
	if err != nil {
		log.Println("数据库信息匹配出错，errmsg:", err)
	}
	err = inifile.Section("service").MapTo(&ConfigInfo.Serconfig)
	if err != nil {
		log.Println("服务信息匹配出错，errmsg:", err)
	}
	err = inifile.Section("redis").MapTo(&ConfigInfo.Redisconfig)
	if err != nil {
		log.Println("redis信息匹配出错，errmsg:", err)
	}

	// 检查host 是smtp是否拼接成string
	err = inifile.Section("email").MapTo(&ConfigInfo.Emailconfig)
	if err != nil {
		log.Println("email信息匹配出错，errmsg:", err)
	}

	fmt.Println(*ConfigInfo)
}
