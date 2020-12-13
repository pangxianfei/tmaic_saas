package users

import (
	"tmaic/app/dao"
	"tmaic/app/models"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct{}

func (user *userService) Get(id int64) *models.User {
	return dao.UserDao.Get(id)
}
func (user *userService) Update(t *models.User) error {
	err := dao.UserDao.Update(t)
	return err
}
