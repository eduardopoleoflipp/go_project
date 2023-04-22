package main

import (
	"fmt"
	"my_project/config"
)

func main() {
	config.InitConfig()
	fmt.Printf("Env is %s, config is %s", config.Env.Name, config.Env.MyConfig)
}
