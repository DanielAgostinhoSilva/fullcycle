package error

import (
	"net/http"
	"time"
)

type Problem struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
}

func NewProblemNotFound(detail string) *Problem {
	return &Problem{
		Status:    http.StatusNotFound,
		Timestamp: time.Now().Local(),
		Title:     "Resource not found",
		Detail:    detail,
	}
}

func NewProblemEntityInUsed(detail string) *Problem {
	return &Problem{
		Status:    http.StatusConflict,
		Timestamp: time.Now().Local(),
		Title:     "Resource is already in use",
		Detail:    detail,
	}
}

func NewProblemBadRequest(detail string) *Problem {
	return &Problem{
		Status:    http.StatusBadRequest,
		Timestamp: time.Now().Local(),
		Title:     "Invalid data",
		Detail:    detail,
	}
}

func NewProblemUnauthorized(detail string) *Problem {
	return &Problem{
		Status:    http.StatusUnauthorized,
		Timestamp: time.Now().Local(),
		Title:     "Unauthorized user",
		Detail:    detail,
	}
}

func NewProblemInternalServerError() *Problem {
	return &Problem{
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now().Local(),
		Title:     "Internal Server Error",
		Detail:    "An unexpected internal error has occurred in the system. Try again and if the problem persists, contact your system administrator.",
	}
}
