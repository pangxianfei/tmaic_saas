package dao

import (
	"tmaic/app/models"
)

var UserDao = newUserDao()

func newUserDao() *userDao {
	return &userDao{}
}

type userDao struct{}

func (r *userDao) Get(id int64) *models.User {

	ret := &models.User{}
	if err := ret.DB().First(ret, "user_id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userDao) Update(t *models.User) (err error) {
	ret := &models.User{
		ID: t.ID,
	}
	err = t.DB().Model(&ret).Save(t).Error
	return
}
