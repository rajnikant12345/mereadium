package main

import (
	"github.com/rajnikant12345/mereadium/app"
	"github.com/rajnikant12345/mereadium/config"
	"github.com/rajnikant12345/mereadium/controller"
)

func main() {
	// initialize service
	c := config.InitConfig()
	app := app.App{}
	app.Initialize(c)
	controller.Initialize(&app)

}
