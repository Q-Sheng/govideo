package gorm

import (
	"gorm.io/gorm"
	"log"
	"service/global"
	"service/model"
)

func InitGorm(configinfo *global.ConfigInformation) *gorm.DB {
	//  连接数据库
	db := GormMysql(configinfo)

	// 初始化创建表格
	migrateTables(db)
	return db
}

// 初始化创建表格
func migrateTables(db *gorm.DB) {
	// 集合模型
	tables := []interface{}{
		model.User{},
	}

	for _, t := range tables {
		err := db.AutoMigrate(&t)
		if err != nil {
			log.Println(err)
		}
	}
}
