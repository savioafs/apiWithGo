package main

import (
	"fmt"
	"log"

	config "github.com/savioafs/apiWithGo/configs"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(cfg.DBDriver)
}
