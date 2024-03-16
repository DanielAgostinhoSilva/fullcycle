package webserver

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	_ "github.com/DanielAgostinhoSilva/fullcycle/9-api/docs"
	pkgMiddleware "github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/webserver/middleware/error"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger/v2"
	"log"
	"sync"
)

func StartServer(wg sync.WaitGroup, configs *enviroment.EnvConfig, controllers ...webserver.Controller) {
	log.Println("Starting web server on port " + configs.WebServerPort)
	defer wg.Done()

	server := webserver.NewWebServer(configs.WebServerPort, configs.TokenAuth)
	server.RegisterMiddleware(middleware.Heartbeat("/health"))
	server.RegisterMiddleware(middleware.Logger)
	server.RegisterMiddleware(pkgMiddleware.ExceptionHandler)
	server.RegisterMiddleware(middleware.WithValue("jwt", configs.TokenAuth))
	server.RegisterMiddleware(middleware.WithValue("jwExpiresIn", configs.JwExpiresIn))
	for _, controller := range controllers {
		server.AddController(controller)
	}
	server.AddHandler("GET", "/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	server.Start()
}
