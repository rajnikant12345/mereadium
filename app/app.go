package app

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rajnikant12345/mereadium/config"
	"github.com/rajnikant12345/mereadium/models"
)

type App struct {
	DB   *gorm.DB
	Conf *config.Config
}

func (a *App) Initialize(c *config.Config) {
	a.SetConfig(c)
	a.InitDatabase()
}

func (a *App) SetConfig(c *config.Config) {
	a.Conf = c
}

func (a *App) InitDatabase() {
	db, err := gorm.Open("postgres", a.Conf.DATABAE_URL)
	if err != nil {
		log.Panicln(err.Error())
	}
	if db == nil {
		panic("Database initialization failed")
	}
	db = db.AutoMigrate(&models.Article{})

	if db == nil {
		panic("Database initialization failed")
	}

	ma := models.Article{}
	ma.ArticleID = "asdf12345"
	ma.UserID = "rajni"
	ma.Title = "Useless article"
	db.Save(&ma)

	a.DB = db
}
