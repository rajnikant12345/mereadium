package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (c *Controller) getArticleList(context echo.Context) error {
	ar := c.Model.GetArticleList()
	return context.JSON(http.StatusOK, &ar)
}

func (c *Controller) AddRoutes() {
	c.echo.GET("/article", c.getArticleList)
}
