package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeohiveConfig()
	fmt.Println("Starting NeoHive with DB:", cfg.DBUrl)
}

