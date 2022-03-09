package requests

// TenantDBRegister 注册验证器
type TenantDBRegister struct {
	TenantId int64 `json:"tenant_id" binding:"required"`
}
