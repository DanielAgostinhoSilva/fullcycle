package main

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs"
	"sync"
)

var (
	cfg *configs.EnvConfig
	db  *sql.DB
)

func init() {
	cfg = configs.LoadConfig("cmd/server/.env")
	configs.MigrationUP(cfg)
	db = configs.DatabaseInitialize(cfg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go configs.WebServerInitialize(wg, cfg, db)
	wg.Wait()
}
