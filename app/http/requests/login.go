package requests

// UserLogin 登陆验证器
type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}
