package models

import (
	"errors"
	"fmt"
	"math"

	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)

const AFFILIATION_CODE_LENGTH uint = 6

type UserAffiliation struct {
	UserID    *int64     `gorm:"column:user_id;primary_key;type:int unsigned"`
	Code      *string    `gorm:"column:uaff_code;type:varchar(32);unique_index;not null"`
	FromCode  *string    `gorm:"column:uaff_from_code;type:varchar(32)"`
	Root      *uint      `gorm:"column:uaff_root_id;type:int unsigned"`
	Parent    *uint      `gorm:"column:uaff_parent_id;type:int unsigned"`
	Left      *uint      `gorm:"column:uaff_left_id;type:int unsigned;not null"`
	Right     *uint      `gorm:"column:uaff_right_id;type:int unsigned;not null"`
	Level     *uint      `gorm:"column:uaff_level;type:int unsigned;not null"`
	CreatedAt *zone.Time `gorm:"column:user_created_at"`
	UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (uaff *UserAffiliation) TableName() string {
	return uaff.SetTableName("user_affiliation")
}

func (uaff *UserAffiliation) Default() interface{} {
	return UserAffiliation{}
}

func (uaff *UserAffiliation) generateCode(user *User) (string, error) {
	if float64(user.ID) > math.Pow(16, float64(AFFILIATION_CODE_LENGTH)) {
		return "", errors.New("userID excceed the max affiliation length")
	}

	affiliationLen := fmt.Sprintf("%d", AFFILIATION_CODE_LENGTH)
	return fmt.Sprintf("%0"+affiliationLen+"x", user.ID), nil
}
