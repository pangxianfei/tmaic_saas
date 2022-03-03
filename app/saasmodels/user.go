package saasmodels

import (
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/saas"
)

type SaasUser struct {
	ID        int64     `gorm:"column:user_id;primary_key;auto_increment"`
	Name      string    `gorm:"column:user_name;type:varchar(100)"`
	TenantsId int64     `gorm:"column:tenants_id;type:int unsigned"`
	Email     string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
	Password  string    `gorm:"column:user_password;type:varchar(100);not null;" json:"-"`
	CreatedAt zone.Time `gorm:"column:user_created_at"`
	UpdatedAt zone.Time `gorm:"column:user_updated_at"`
	DeletedAt zone.Time `gorm:"column:user_deleted_at"`
	saas.BaseModel
}
