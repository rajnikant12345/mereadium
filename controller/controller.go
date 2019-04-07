package controller

import (
	"log"

	"github.com/labstack/echo"

	"github.com/rajnikant12345/mereadium/config"
	"github.com/rajnikant12345/mereadium/models"
)

type Controller struct {
	Config *config.Config
	Model  *models.Model
	echo   *echo.Echo
}

func NewController(config *config.Config, model *models.Model) *Controller {
	controller := Controller{}
	controller.Config = config
	controller.Model = model
	controller.echo = echo.New()
	return &controller
}

func (c *Controller) Serve() {
	c.AddRoutes()
	log.Println("Starting the server...")
	c.echo.Logger.Fatal(c.echo.Start(":13230"))
}
