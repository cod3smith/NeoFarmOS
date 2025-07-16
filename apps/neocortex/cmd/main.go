package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeocortexConfig()
	fmt.Println("Starting NeoCortex with DB:", cfg.DBUrl)
}

