package main

import (
	"fmt"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs"
)

func main() {
	cfg := configs.LoadConfig("cmd/server/.env")
	fmt.Println(cfg.DBName)
}
