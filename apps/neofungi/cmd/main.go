package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeofungiConfig()
	fmt.Println("Starting NeoFungi with DB:", cfg.DBUrl)
}

