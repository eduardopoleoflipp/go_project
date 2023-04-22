package main

import (
	"fmt"
	"my_project/config"
)

func main() {
	config.InitConfig()
	fmt.Printf(
		"Env is %s\n MyConfig is %s\n DefaultConfig is %s\n NotSetConfig, %s\n",
		config.AppConfig.Name,
		config.AppConfig.MyConfig,
		config.AppConfig.DefaultConfig,
		config.AppConfig.NotSetConfig,
	)
}
