package usersRou

import (
	"github.com/gin-gonic/gin"
	v1 "service/api/v1"
)

type LoginRouter struct {
}

func (this *LoginRouter) InitLoginRouter(routerGro *gin.RouterGroup) {
	loginRouter := routerGro.Group("login").Use()
	userapi := v1.APIGroupApp.Userapi
	{
		loginRouter.POST("/register", userapi.Register)
	}

}
