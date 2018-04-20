package model

import "net/http"

type Error struct {
	Code    int
	Message string
}

func GetError(code int, message string) Error {
	return Error{code, message}
}

func GetAccessDeniedError() Error {
	return Error{http.StatusForbidden, "Access Denied"}
}

func GetInternalServerErrorError() Error {
	return Error{http.StatusInternalServerError, "Internal server error"}
}
