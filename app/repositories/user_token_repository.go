package repositories

import (
	"gorm.io/gorm"

	sysmdel "gitee.com/pangxianfei/frame/model/sysmodel"
)

var UserTokenRepository = newUserTokenRepository()

func newUserTokenRepository() *userTokenRepository {
	return &userTokenRepository{}
}

type userTokenRepository struct {
}

func (r *userTokenRepository) GetByToken(db *gorm.DB, token string) *sysmdel.UserToken {
	if len(token) == 0 {
		return nil
	}
	return r.Take(db, "token = ?", token)
}

func (r *userTokenRepository) Get(db *gorm.DB, id int64) *sysmdel.UserToken {
	ret := &sysmdel.UserToken{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.UserToken {
	ret := &sysmdel.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Create(db *gorm.DB, t *sysmdel.UserToken) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userTokenRepository) Update(db *gorm.DB, t *sysmdel.UserToken) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userTokenRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&sysmdel.UserToken{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userTokenRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&sysmdel.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userTokenRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&sysmdel.UserToken{}, "id = ?", id)
}
