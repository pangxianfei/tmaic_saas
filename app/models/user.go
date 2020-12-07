package models

import (
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/helpers/ptr"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)

type User struct {
	model.BaseModel
	ID        *uint      `gorm:"column:user_id;primary_key;auto_increment"`
	Name      *string    `gorm:"column:user_name;type:varchar(100)"`
	Email     *string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
	Password  *string    `gorm:"column:user_password;type:varchar(100);not null"`
	CreatedAt *zone.Time `gorm:"column:user_created_at"`
	UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
}

func (user *User) TableName() string {
	return user.SetTableName("user")
}
func (user *User) Default() interface{} {
	return User{}
}
func (user *User) Scan(userId uint) error {
	newUser := User{
		ID: ptr.Uint(userId),
	}
	if err := m.H().First(&newUser, false); err != nil {
		return err
	}
	*user = newUser
	return nil
}
func (user *User) Value() interface{} {
	return user
}

func (user *User) User() *User {
	return user
}

func (user *User) SetNameAttribute(value interface{}) {
	user.Name = user.Email
}

func (user *User) GetPasswordAttribute(value interface{}) interface{} {
	return value
}
