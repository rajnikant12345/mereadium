package app

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rajnikant12345/mereadium/models"
	"net/http"
	"net/url"
)



func (a *App) GetArticles(values url.Values) []models.Article {
	m := []models.Article{}
	a.DB.Find(&m)
	return m
}

func (a *App) DeleteArticle(id string)  error {
	uid,err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}
	tmp := models.Article{}
	tmp.ArticleID = uid
	db := a.DB.Where("article_id = ?", tmp.ArticleID).Delete(&models.Article{})
	if db.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound,db.Error.Error())
	}
	return nil
}


func (a *App) GetArticle(id string) (*models.Article, error) {
	m := models.Article{}
	uid,err := uuid.Parse(id)
	if err != nil {
		return nil,echo.NewHTTPError(http.StatusNotFound,err.Error())
	}
	tmp := models.Article{}
	tmp.ArticleID = uid

	db := a.DB.Where("article_id = ?", tmp.ArticleID).First(&m)
	if db.Error != nil {
		return nil,echo.NewHTTPError(http.StatusNotFound,db.Error.Error())
	}
	return &m,nil
}



func (a *App) PostArticle(context echo.Context) (*models.Article,error) {
	m := models.Article{}
	err := context.Bind(&m)
	if err != nil {
		return nil,echo.NewHTTPError(http.StatusBadRequest, err.Error() )
	}

	db := a.DB.Create(&m)
	if db.Error != nil {
		return nil,echo.NewHTTPError(http.StatusBadRequest, db.Error )
	}

	return &m,nil

}
