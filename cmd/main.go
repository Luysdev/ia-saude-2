package main

import (
	"saude/config"
	"saude/internal/router"
)

func main() {
	config.InitConfig()
	router.InitRouter()
}
