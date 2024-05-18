package servicegroup

import (
	"log"
	"service/global"
	"service/model"
)

// 与user有关的业务，具体的数据库操作, 返回给接口用

type UserService struct{}

// 添加数据
func (us *UserService) Register(u model.User) (model.User, error) {
	// todo 查询用户命名是否存在
	// return nil, errors.New("用户名已注册")

	// todo 对u的密码进行加密
	err := global.Db.Create(&u).Error
	if err != nil {
		log.Println("Create失败", err)
	}
	return u, err
}
