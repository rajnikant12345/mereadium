package controller

import (
	"github.com/labstack/echo"
	"github.com/rajnikant12345/mereadium/app"
	"log"
	"net/http"
)

func GetArticles(a *app.App) func(context echo.Context) error {
	return func(context echo.Context) error {
		query := context.QueryParams()
		articles := a.GetArticles(query)
		return context.JSON(http.StatusOK , &articles)
	}
}

func Initialize(app *app.App) {
	e := echo.New()
	e.GET("/articles",GetArticles(app))
	log.Println("Starting the server...")
	e.Logger.Fatal(e.Start(":" +app.Conf.PORT))

}
