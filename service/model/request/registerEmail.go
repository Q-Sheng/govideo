package request

// 注册时上下文传入的email
// binding:"required,email"表示email字段是必需的，并且必须是一个有效的电子邮件地址
type RegisterEmail struct {
	Email string `json:"email" binding:"required,email"`
}
