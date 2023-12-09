package error

import (
	"encoding/json"
	"fmt"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
	"github.com/go-errors/errors"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer handler(w, r)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err, ok := recover().(error); ok {
		switch err.(type) {
		case *exception.EntityNotFoundError:
			handlerError(NewProblemNotFound(err.Error()), http.StatusNotFound, w)
		case *exception.EntityInUsedError:
			handlerError(NewProblemEntityInUsed(err.Error()), http.StatusConflict, w)
		default:
			handlerError(NewInternalServerError(), http.StatusInternalServerError, w)
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
