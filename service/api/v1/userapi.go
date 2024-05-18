package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service/model"
)

type UserApi struct{}

// 根据contex 添加 用户
func (ua *UserApi) Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println("绑定失败：", err)
	}

	// todo 验证

	// 调用业务层
	userReturn, err := userService.Register(user)
	if err != nil {
		log.Println("业务层出错：", err)
	}
	// 返回信息给前端
	// todo 可集成为respone

	c.JSON(http.StatusOK, gin.H{
		"data": userReturn,
	})
}
