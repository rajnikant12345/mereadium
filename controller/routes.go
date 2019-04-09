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


func PostArticle(a *app.App) func(context echo.Context) error {
	return func(context echo.Context) error {
		articles,err := a.PostArticle(context)
		if err != nil {
			return err
		}
		return context.JSON(http.StatusOK , &articles)
	}
}


func GetArticle(a *app.App) func(context echo.Context) error {
	return func(context echo.Context) error {
		m,err := a.GetArticle(context.Param("id"))
		if err != nil {
			return err
		}
		return context.JSON(http.StatusOK , m)
	}
}


func DeleteArticle(a *app.App) func(context echo.Context) error {
	return func(context echo.Context) error {
		err := a.DeleteArticle(context.Param("id"))
		if err != nil {
			return err
		}
		return context.String(http.StatusNoContent, "DELETED")
	}
}



func Initialize(app *app.App) {
	e := echo.New()

	e.GET("/article",GetArticles(app))
	e.GET("/article/:id",GetArticle(app))
	e.POST("/article",PostArticle(app))
	e.DELETE("/article/:id",DeleteArticle(app))
	log.Println("Starting the server...")
	e.Logger.Fatal(e.Start(":" +app.Conf.PORT))

}
