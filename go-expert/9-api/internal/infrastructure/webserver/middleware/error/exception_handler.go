package error

import (
	"encoding/json"
	"fmt"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
	webserver "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver/exception"
	"github.com/go-errors/errors"
	"net/http"
)

func ExceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer handler(w, r)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err, ok := recover().(error); ok {
		var entityNotFoundError *exception.EntityNotFoundError
		var entityInUsedError *exception.EntityInUsedError
		var badRequestError *webserver.BadRequestError
		var unauthorized *webserver.UnauthorizedError
		var unauthorizedUser *exception.UnauthorizedUserError
		switch {
		case errors.As(err, &entityNotFoundError):
			handlerError(NewProblemNotFound(err.Error()), http.StatusNotFound, w)
		case errors.As(err, &entityInUsedError):
			handlerError(NewProblemEntityInUsed(err.Error()), http.StatusConflict, w)
		case errors.As(err, &badRequestError):
			handlerError(NewProblemBadRequest(err.Error()), http.StatusBadRequest, w)
		case errors.As(err, &unauthorized):
			handlerError(NewProblemUnauthorized(err.Error()), http.StatusUnauthorized, w)
		case errors.As(err, &unauthorizedUser):
			handlerError(NewProblemUnauthorized(err.Error()), http.StatusUnauthorized, w)
		default:
			handlerError(NewProblemInternalServerError(), http.StatusInternalServerError, w)
		}
		logError(err)
	}

}

func handlerError(problem *Problem, httpStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(problem)
}

func logError(err interface{}) {
	wrappedErr := errors.Wrap(err, 0)
	fmt.Println(wrappedErr.ErrorStack())
}
