package router

import (
	"service/router/usersRou"
)

type RouterGroup struct {
	usersRou.UserRouter
}

var RouterGroupApp = new(RouterGroup)

//
//func InitRouter() {
//
//	router := gin.Default()
//
//	//跨域
//
//	//路由组
//	// PublicRouter: 公共路由组
//
//	// PrivateRouter:
//	PublicRouter := router.Group("public")
//	{
//		PublicRouter.GET("/test", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"status":  http.StatusOK,
//				"message": "test public",
//			})
//		})
//	}
//	//PrivateRouter := router.Group("private")
//	//{
//	//	//初始化路由组
//	//
//	//}
//
//	err := router.Run()
//	if err != nil {
//		fmt.Printf("errmsg:", err)
//	}
//
//	return
//
//}
