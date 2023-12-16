package main

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/migration"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/webserver"
	"sync"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wesley Willians
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := enviroment.LoadConfig("cmd/server/.env")
	migration.MigrationUP(cfg)
	db := database.InitializeDatabase(cfg)
	defer db.Close()
	productController := InitializeProductController(db)
	userController := InitializeUserController(db)

	var wg sync.WaitGroup
	wg.Add(1)
	go webserver.StartServer(wg, cfg, productController, userController)
	wg.Wait()
}
