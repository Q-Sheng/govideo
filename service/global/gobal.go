package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

/*
	全局变量
*/

// 解开import循环
// import (
//
//	"service/initialize"
//
// )
// // gobal的部分全局变量，由initialize中的全局变量提供
// var (
//
//	Db         = initialize.Db
//	ConfigInfo = initialize.ConfigInfo
//	Router     = initialize.Router
//
// gobal的部分全局变量，由initialize中的全局变量提供
var (
	Db         = new(gorm.DB)
	Rdb        = new(redis.Client)
	ConfigInfo = new(ConfigInformation)
	Router     = new(gin.Engine)
)
