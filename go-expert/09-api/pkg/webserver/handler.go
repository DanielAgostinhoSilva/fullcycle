package webserver

import "net/http"

type Handler struct {
	Method      string
	HandlerFunc http.HandlerFunc
}
