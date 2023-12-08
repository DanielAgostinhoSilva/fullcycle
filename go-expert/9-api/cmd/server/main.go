package main

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/migration"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/webserver"
	"sync"
)

func main() {
	cfg := enviroment.LoadConfig("cmd/server/.env")
	migration.MigrationUP(cfg)
	db := database.InitializeDatabase(cfg)
	defer db.Close()
	productController := InitializeProductController(db)

	var wg sync.WaitGroup
	wg.Add(1)
	go webserver.StartServer(wg, cfg, productController)
	wg.Wait()
}
