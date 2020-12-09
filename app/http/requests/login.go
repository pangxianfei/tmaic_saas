package requests

type UserLogin struct {
	ID       int64  `json:"user_id"`
	Name     string `json:"user_name;"`
	Email    string `json:"user_email;"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}
