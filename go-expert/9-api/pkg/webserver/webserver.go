package webserver

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	Controllers   []Controller
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		Controllers:   []Controller{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddController(controller Controller) {
	s.Controllers = append(s.Controllers, controller)
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = Handler{Method: method, HandlerFunc: handler}
}

// register middleware
func (s *WebServer) RegisterMiddleware(middleware func(http.Handler) http.Handler) {
	s.Router.Use(middleware)
}

// Start loop through the handlers and add them to the router
// start the server
func (s *WebServer) Start() {
	for _, controller := range s.Controllers {
		s.Router.Route(controller.Path(), controller.Router)
	}
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.HandlerFunc)
	}
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
