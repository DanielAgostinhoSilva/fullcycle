package main

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs"
)

func main() {
	cfg := configs.LoadConfig("cmd/server/.env")
	configs.MigrationUP(cfg)
}
