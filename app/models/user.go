package models

import (
	"gitee.com/pangxianfei/frame/helpers/hash"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/helpers/ptr"
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/model"
)

type User struct {
	ID        int64      `gorm:"column:user_id;primary_key;auto_increment"`
	Name      *string    `gorm:"column:user_name;type:varchar(100)"`
	TenantsId int64      `gorm:"column:tenants_id;type:int unsigned"`
	Email     *string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
	Password  *string    `gorm:"column:user_password;type:varchar(100);not null;" json:"-"`
	CreatedAt *zone.Time `gorm:"column:user_created_at"`
	UpdatedAt *zone.Time `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (user *User) TableName() string {
	return user.SetTableName("user")
}

func (user *User) Default() interface{} {
	return User{}
}

func (user *User) Scan(userId int64) error {
	newUser := User{
		ID: userId,
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
	// model.DB().Where("user_id = ?", 1).Find(user)
	return user
}

// SetNameAttribute Mutator: auto generate user name
func (user *User) SetNameAttribute(value interface{}) {
	user.Name = ptr.String(hash.Md5(*user.Email))
}

func (user *User) ObjArr(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (interface{}, error) {
	var outArr []User
	if err := m.H().Q(filterArr, sortArr, limit, withTrashed).Find(&outArr).Error; err != nil {
		return nil, err
	}
	return outArr, nil
}
func (user *User) ObjArrPaginate(c model.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []User
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, int64(perPage))
}
