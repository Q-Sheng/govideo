package v1

import "service/servicegroup"

/*
Application Programming Interface:程序之间的接口。程序之间的合约。
沟通那个router与业务逻辑层
*/
type ApiGroup struct {
	Userapi UserApi
}

// 声明Api结合供router调用
var APIGroupApp = new(ApiGroup)

var (
	userService = servicegroup.UserService{}
)
