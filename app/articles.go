package app

import (
	"github.com/rajnikant12345/mereadium/models"
	"net/url"
)



func (a *App) GetArticles(values url.Values) []models.Article {
	m := []models.Article{}
	a.DB.Find(&m)
	return m
}
