package usersRou

import (
	"github.com/gin-gonic/gin"
	v1 "service/api/v1"
)

type LoginRouter struct {
}

// InitLoginRouter
func (r *LoginRouter) InitLoginRouter(routerGro *gin.RouterGroup) {
	loginRouter := routerGro.Group("login").Use()
	userapi := v1.APIGroupApp.Userapi
	{
		// todo 注册没做对用户检查是否已存在。同时需要全部关键数据不为空再对数据库操作
		loginRouter.POST("/register", userapi.Register)
		// todo about email
		// send verification code by email
		loginRouter.POST("/login/sendEmailVerificationCode", userapi.SendEmailVerCode)

		// 登录
		loginRouter.POST("/login", userapi.LoginUser)
		// 获取用户信息
		//loginRouter.POST("/getUserInfo", )
		// 修改用户信息
		//loginRouter.POST("/updateUser",)

	}
}
