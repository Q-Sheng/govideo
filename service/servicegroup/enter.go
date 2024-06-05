package servicegroup

// service层，业务逻辑问题

type ServiceGroup struct {
	Userservice UserService //匿名字段，将UserService的方法给到ServiceGroup使用
}

// 声明service拱api调用
var SerGroup = new(ServiceGroup)
