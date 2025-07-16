package main

import (
	"fmt"

	"github.com/cod3smith/NeoFarmOS/shared/config"
)

func main() {
	cfg := config.GetNeobiolabConfig()
	fmt.Println("Starting NeoBioLab with DB:", cfg.DBUrl)
}

