package config

import "os"

type Config struct {
	PORT        string
	DATABAE_URL string
}

func InitConfig() *Config {
	c := Config{}
	c.DATABAE_URL = os.Getenv("DATABASE_URL")
	c.PORT = os.Getenv("PORT")
	return &c
}
