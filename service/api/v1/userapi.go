package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service/model"
	"service/model/request"
	"service/model/response"
	service "service/servicegroup"
)

type UserApi struct{}

// 根据contex 添加 用户
func (ua *UserApi) Register(c *gin.Context) {
	var ru model.RegisterUser
	err := c.ShouldBindJSON(&ru)

	if err != nil {
		log.Println("绑定失败：", err)
	}

	// todo 验证

	// 调用业务层
	// todo 是否返回user类型
	err = service.SerGroup.Userservice.Register(&ru)
	if err != nil {
		log.Println("业务层出错：", err)
	}
	// 返回信息给前端
	if err != nil {
		response.ResFail(c, err)
	} else {
		response.ResSuccess(c, http.StatusOK, "success", ru)
	}
}

// 登录
func (ua *UserApi) LoginUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println("绑定失败：", err)
	}

	// todo 验证

	// 调用业务层
	err = service.SerGroup.Userservice.Login(&user)
	if err != nil {
		log.Println("业务层出错：", err)
	}
	// 返回信息给前端, 集成为response
	if err != nil {
		response.ResFail(c, err)
	} else {
		response.ResSuccess(c, http.StatusOK, "success", user)
	}
}

// 修改个人信息，前端无传入ID，该如何
func (ua *UserApi) SetUserInfo(c *gin.Context) {
	// 声明一个来提取上下文的信息
	var user model.User
	id := c.GetUint("uid")
	user.ID = id
	// 检查是否有ID，是否需要自建结构体绑定json
	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println("Binding JSON Failed")
	}

	//调用业务层 应该不用Userreturn了吧，传送了地址变量
	err = service.SerGroup.Userservice.SetInfo(&user)

	if err != nil {
		response.ResFail(c, err)
	} else {
		response.ResSuccess(c, http.StatusOK, "success", user)
	}
}

// todo 注册，发送验证到邮箱
func (ua *UserApi) SendEmailVerCode(c *gin.Context) {

	//调用service，在什么条件下？
	//满足条件：注册情况下，邮箱未被使用。所以需要先json绑定邮箱
	var registeremail request.RegisterEmail
	err := c.ShouldBindJSON(&registeremail)
	if err != nil {
		log.Println("binding failed:", err)
	} else {
		//进入业务层
		err = service.SerGroup.Userservice.SendEmailCode(&registeremail)
	}

	if err != nil {
		response.ResFail(c, err)
	} else {
		// 直接返回给前端来验证不妥吧
		response.ResSuccess(c, http.StatusOK, "success", "发送成功")
	}

}
