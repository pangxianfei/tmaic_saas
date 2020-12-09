package models

import (
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)

type Article struct {
	model.BaseModel
	ID        uint       `gorm:"column:id;primary_key;auto_increment"`
	Title     string     `gorm:"column:title;type:varchar(100)"`
	Body      string     `gorm:"column:body;type:varchar(100);unique_index;not null"`
	Slug      string     `gorm:"column:slug;type:varchar(11);unique_index;not null"`
	CreatedAt zone.Time  `gorm:"column:created_at"`
	UpdatedAt zone.Time  `gorm:"column:updated_at"`
	DeletedAt *zone.Time `gorm:"column:deleted_at"`
}

func (Article *Article) TableName() string {
	return Article.SetTableName("article")
}

func (Article *Article) SetSlugAttribute(value interface{}) {
	Article.Slug = Article.Title
}
