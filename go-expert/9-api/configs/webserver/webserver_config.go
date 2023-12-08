package webserver

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	pkgMiddleware "github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/webserver/middleware/error"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"sync"
)

func StartServer(wg sync.WaitGroup, configs *enviroment.EnvConfig, controllers ...webserver.Controller) {
	log.Println("Starting web server on port " + configs.WebServerPort)
	defer wg.Done()

	server := webserver.NewWebServer(configs.WebServerPort)
	server.RegisterMiddleware(middleware.Logger)
	server.RegisterMiddleware(pkgMiddleware.ErrorHandler)
	for _, controller := range controllers {
		server.AddController(controller)
	}
	server.Start()
}
