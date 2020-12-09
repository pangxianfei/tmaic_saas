package migrations

import (
	"gorm.io/gorm"

	"github.com/pangxianfei/framework/database/migration"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)

func init() {
	migration.AddMigrator(&CreateFailedQueueTable1556612225{})
}

type FailedQueue struct {
	ID *uint `gorm:"column:failed_queue_id;primary_key;auto_increment"`

	Hash      *string        `gorm:"column:failed_queue_hash;type:varchar(100);unique_index;not null"`
	Topic     *string        `gorm:"column:failed_queue_topic_name;type:varchar(100);not null"`
	Channel   *string        `gorm:"column:failed_queue_channel_name;type:varchar(100);not null"`
	DataProto *[]byte        `gorm:"column:failed_queue_data;type:varbinary(2048)"`
	PushedAt  *zone.Time     `gorm:"column:failed_queue_pushed_at;not null"`
	Delay     *zone.Duration `gorm:"column:failed_queue_delay;type:bigint unsigned;not null"`
	Retries   *uint32        `gorm:"column:failed_queue_retries;type:integer unsigned;not null"`
	Tried     *uint32        `gorm:"column:failed_queue_tried;type:integer unsigned;not null"`
	Err       *string        `gorm:"column:failed_queue_err;type:longtext"`

	CreatedAt *zone.Time `gorm:"column:failed_queue_created_at"`
	UpdatedAt zone.Time  `gorm:"column:failed_queue_updated_at"`
	DeletedAt *zone.Time `gorm:"column:failed_queue_deleted_at"`
	model.BaseModel
}

func (fq *FailedQueue) TableName() string {
	return fq.SetTableName("failed_queue")
}

type CreateFailedQueueTable1556612225 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*CreateFailedQueueTable1556612225) Up(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().CreateTable(&FailedQueue{})
	return db
}

func (*CreateFailedQueueTable1556612225) Down(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().DropTable(&FailedQueue{})
	return db
}
