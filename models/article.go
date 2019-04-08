package models

import "github.com/jinzhu/gorm/dialects/postgres"

type Article struct {
	ArticleID string         `json:"article_id" , gorm:"primary_key`
	UserID    string         `json:"user_id" , gorm:"not null"`
	Title     string         `json:"title" , gorm:"not null"`
	Meta      postgres.Jsonb `json:"tags , omitempty"`
	URI       string         `json:"uri"`
	Comments  postgres.Jsonb `json:"comments,omitempty"`
}

func GetArticleList() []Article {
	return nil
}
