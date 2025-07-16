package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeosilkConfig()
	fmt.Println("Starting NeoSilk with DB:", cfg.DBUrl)
}

