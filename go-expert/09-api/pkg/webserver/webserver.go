package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	Controllers   []Controller
	WebServerPort string
	TokenAuth     *jwtauth.JWTAuth
}

func NewWebServer(serverPort string, tokenAuth *jwtauth.JWTAuth) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		Controllers:   []Controller{},
		WebServerPort: serverPort,
		TokenAuth:     tokenAuth,
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
		s.Router.Route(controller.Path(), controller.Router(s.TokenAuth))
	}
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.HandlerFunc)
	}
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
