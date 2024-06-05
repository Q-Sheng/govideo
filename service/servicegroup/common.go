package servicegroup

import (
	"log"
	"service/global"
)

// 根据字段检查用户中是否存在数据，匹配用户名或邮箱
func IsExistField(thisModel interface{}, field string, value any) bool {
	dbModel := global.Db.Where(field, value).Find(&thisModel)

	// thisModel为所要检索的表，ID为主键，若为0则说明没有找到，如何判断是否找到
	if dbModel.Error != nil {
		log.Println(dbModel.Error)
		thisModel = nil
		return false
	}
	if dbModel.RowsAffected <= 0 {
		thisModel = nil
		return false
	}
	return true
}
