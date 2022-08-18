package main

import (
	"housing_viz/pkg/common"
	"housing_viz/pkg/web"
	"os"
)

func main() {
	if os.Getenv("DOCKER") != "TRUE" {
		common.LoadEnv()
	}

	web.InitServer()
}
