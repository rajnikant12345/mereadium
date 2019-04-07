package app

import (
	"github.com/rajnikant12345/mereadium/config"
	"github.com/rajnikant12345/mereadium/controller"
	"github.com/rajnikant12345/mereadium/models"
)

func NewApp() {
	m := models.NewModel()
	conf := config.InitConfig()
	cont := controller.NewController(conf, m)
	cont.Serve()
}
