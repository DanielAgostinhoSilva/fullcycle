package configs

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/webserver/controller"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver"
	pkgMiddleware "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver/middleware/error"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"sync"
)

func WebServerInitialize(wg sync.WaitGroup, configs *EnvConfig, db *sql.DB) {
	log.Println("Starting web server on port " + configs.WebServerPort)
	defer wg.Done()
	productController := controller.NewProductController(database.NewProductRepository(db))

	server := webserver.NewWebServer(configs.WebServerPort)
	server.RegisterMiddleware(middleware.Logger)
	server.RegisterMiddleware(pkgMiddleware.ErrorHandler)
	server.AddController(productController)
	server.Start()
}
