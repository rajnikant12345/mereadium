package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Article struct {
	ArticleID uuid.UUID      `json:"article_id"`
	UserID    string         `json:"user_id"`
	Title     string         `json:"title"`
	Meta      postgres.Jsonb `json:"tags , omitempty"`
	URI       string         `json:"uri"`
	Comments  postgres.Jsonb `json:"comments,omitempty"`
}


func (user *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("article_id",uuid.New())
	return nil
}
