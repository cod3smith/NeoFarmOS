package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeoedgeConfig()
	fmt.Println("Starting NeoEdge with DB:", cfg.DBUrl)
}

