package dao

import (
	"tmaic/app/models"
)

var UserDao = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r *userRepository) Get(id int64) *models.User {

	ret := &models.User{}
	if err := ret.DB().First(ret, "user_id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) Update(t *models.User) (err error) {
	ret := &models.User{
		ID: t.ID,
	}
	err = t.DB().Model(&ret).Save(t).Error
	return
}
