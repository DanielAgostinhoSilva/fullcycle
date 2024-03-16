package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Controller interface {
	Router(tokenAuth *jwtauth.JWTAuth) func(r chi.Router)
	Path() string
}
