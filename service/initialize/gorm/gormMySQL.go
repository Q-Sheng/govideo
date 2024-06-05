package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"service/global"
)

func GormMysql(Dbconfig *global.DbConfig) *gorm.DB {
	// todo 检查configinfo各值不为空

	mysqlConfig := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			Dbconfig.DbUser,
			Dbconfig.DbPassword,
			Dbconfig.DbHost,
			Dbconfig.DbPort,
			Dbconfig.DbName,
		),
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("errmsg:", err)
		return nil
	} else {
		return db
	}

}
