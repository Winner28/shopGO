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
	return Error{http.StatusInternalServerError, "Internal server error."}
}

func GetNotFoundError() Error {
	return Error{http.StatusNotFound, "Not found"}
}

func NotAuthorizedError() Error {
	return Error{http.StatusForbidden, "To visit this page you need to login first"}
}

type Message struct {
	Header  string
	Message string
}

func UserSuccessfullyDeleted() Message {
	return Message{"Deleted", "User successfully deleted."}
}
