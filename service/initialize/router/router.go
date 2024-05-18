package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	myrouter "service/router"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//todo 使用跨域

	// 路由分组
	PublicGroup := router.Group("public")
	{
		// 测试
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// 声明路由变量

	PrivateGroup := router.Group("private")
	{
		// 为路由组添加路由
		myrouter.RouterGroupApp.InitLoginRouter(PrivateGroup)

	}

	return router

}
