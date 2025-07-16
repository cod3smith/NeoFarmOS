package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeoplantConfig()
	fmt.Println("Starting NeoPlant with DB:", cfg.DBUrl)
}

