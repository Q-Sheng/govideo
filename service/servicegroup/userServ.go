package servicegroup

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"service/global"
	"service/model"
	"service/model/request"
	"service/servicegroup/consts"
	"time"
)

// 与user有关的业务，具体的数据库操作, 返回给接口用

type UserService struct{}

// 注册
func (us *UserService) Register(ru *model.RegisterUser) error {
	var u model.User
	if IsExistField(&u, "email", ru.Email) {
		return fmt.Errorf("this email already registered")
	}
	//
	vercode := fmt.Sprintf("%s%s", consts.RegEmailVerCode, ru.Email)
	val, err := global.Rdb.Get(context.Background(), vercode).Result()
	if ru.VerCode != val {
		return fmt.Errorf("验证码错误")
	}

	// todo 查询用户命名否存在
	if IsExistField(&u, "username", ru.Username) {
		return fmt.Errorf("用户名已存在")
	}

	// 赋值给u
	{
		u.Username = ru.Username
		//todo 对u的密码进行加密，赋盐值
		u.Password = ru.Password
		u.Email = ru.Email
	}

	err = global.Db.Create(&u).Error
	if err != nil {
		log.Println("Create失败", err)
	}
	return err
}

// 登录
func (us *UserService) Login(u *model.User) error {
	// todo 查询用户命名是否存在
	var user model.User
	if isExist := IsExistField(&user, "username", u.Username); !isExist {
		return fmt.Errorf("账户不存在") // 错误返回nil，不暴露用户信息
	} else {
		// todo 对u的密码进行加密
		if !us.IsCorrectPass(user.Password, u.Password) {
			return fmt.Errorf("密码错误") // 错误返回nil，不暴露用户信息
		}
		// 将user给u
		*u = user

		return nil // 正确则返回全部用户信息

	}

}

// 注册发送邮箱验证码
func (us *UserService) SendEmailCode(regemail *request.RegisterEmail) error {
	// 查找邮箱是否已存在
	var user model.User
	if isExist := IsExistField(&user, "email", regemail.Email); isExist {
		return fmt.Errorf("邮箱已注册")
	}

	mailTo := []string{regemail.Email}
	// 邮件主题
	title := "验证码"
	// 生成6为随机验证码
	verCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	// body
	body := fmt.Sprintf("您的注册验证码为：%s, 5分钟内有效，请勿转发他人", verCode)
	// 使用gomail进行发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", global.ConfigInfo.Emailconfig.User)
	// 收件
	m.SetHeader("To", mailTo...)
	m.SetHeader("Title", title)
	m.SetBody("text/html", body)

	//
	d := gomail.NewDialer(global.ConfigInfo.Emailconfig.Host, global.ConfigInfo.Emailconfig.Port,
		global.ConfigInfo.Emailconfig.User, global.ConfigInfo.Emailconfig.Pass)
	err := d.DialAndSend(m)

	if err != nil {
		return fmt.Errorf("发送邮件失败。err:", err)
	}
	// 果然设置全局变量，使用redis，同时设置时间
	// 设置验证码到redis中
	err = global.Rdb.Set(context.Background(), fmt.Sprintf("%s%s", consts.RegEmailVerCode, regemail.Email), verCode, 5*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("存储code到redis出错", err)
	}

	return err
}

// todo 根据传入的user信息设置update个人信息，再修改数据库，如何获取user的ID
func (us *UserService) SetInfo(u *model.User) error {
	// 根据uid进行查询

	update := map[string]interface{}{
		"Username": u.Username,
		"Gender":   u.Gender, //是否引用时自动绑定gorm
	}

	// 根据传入user数据包装进update
	// 定位user表,传入u还是应该新定义一个User类型的user
	err := global.Db.Model(u).Where("id", u.ID).Updates(update).Error
	if err != nil {
		log.Println("更新失败，", err)
	}

	return err
}

// 包装加密部分
func (us *UserService) IsCorrectPass(uPass string, password string) bool {
	if uPass == password {
		return true
	}
	return false
}
