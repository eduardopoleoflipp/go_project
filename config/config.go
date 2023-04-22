package config

import "os"

type Config struct {
	Name     string
	MyConfig string
}

var Env Config

func InitConfig() Config {
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "development":
		Env = development
	case "staging":
		Env = staging
	case "production":
		Env = production
	default:
		panic("Environment not setup")
	}

	return Env
}
