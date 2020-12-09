package migrations

import (
	"gorm.io/gorm"

	"github.com/pangxianfei/framework/database/migration"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)

func init() {
	migration.AddMigrator(&CreateUserTable1548750742{})
}

type User struct {
	ID        *uint      `gorm:"column:user_id;primary_key;auto_increment"`
	Name      *string    `gorm:"column:user_name;type:varchar(100)"` //@cautions struct member must be pointer when member could be null
	Email     *string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
	Password  *string    `gorm:"column:user_password;type:varchar(100);not null"`
	CreatedAt *zone.Time `gorm:"column:user_created_at"`
	UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (u *User) TableName() string {
	return u.SetTableName("user")
}

type CreateUserTable1548750742 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*CreateUserTable1548750742) Up(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().CreateTable(&User{})
	return db
}

func (*CreateUserTable1548750742) Down(db *gorm.DB) *gorm.DB {

	_ = db.Migrator().DropTable(&User{})

	return db
}
