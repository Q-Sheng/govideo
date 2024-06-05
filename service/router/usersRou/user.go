package usersRou

import (
	"github.com/gin-gonic/gin"
	v1 "service/api/v1"
)

// 增加个人user路由组
type MyUserRouter struct {
}

// 修改用户信息:/user/setUserInfo
func (u *MyUserRouter) InitUserRouter(routerGro *gin.RouterGroup) {
	// 创建使用路由组
	userRouter := routerGro.Group("user").Use()
	// 引入api
	// todo 如何获取用户的iD
	useapi := v1.APIGroupApp.Userapi
	userRouter.POST("/setUserInfo", useapi.SetUserInfo)

}
