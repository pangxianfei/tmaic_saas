package migrations

import (
	"gorm.io/gorm"

	"gitee.com/pangxianfei/frame/database/migration"
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/model"
)

func init() {
	migration.AddMigrator(&CreateUserAffiliationTable1553678539{})
}

type UserAffiliation struct {
	UserID    *uint      `gorm:"column:user_id;primary_key;type:int unsigned"`
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

type CreateUserAffiliationTable1553678539 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*CreateUserAffiliationTable1553678539) Up(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().CreateTable(&UserAffiliation{})
	return db
}

func (*CreateUserAffiliationTable1553678539) Down(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().DropTable(&UserAffiliation{})
	return db
}
