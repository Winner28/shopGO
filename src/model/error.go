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
	BackTo  string
}

func GetDefaultMessage(header, message, backTo string) Message {
	return Message{header, message, backTo}
}

func UserSuccessfullyDeleted() Message {
	return Message{"Deleted", "User successfully deleted.", "users"}
}

func WeDontHaveSuchUser() Message {
	return Message{"Look like you got lost", "We dont have such User in our system", "users"}
}

func ErrorWhileUpdatingUser(message string) Message {
	return Message{"System problems", message, "users"}
}

func ErrorWhileCreatingUser(message string) Message {
	return Message{"System problems", message, "users"}

}

func UserSuccessfullyUpdated(user User) Message {
	return Message{"Successfully updated", "User with EMAIL: " + user.Email + " updated", "users"}
}

func UserSuccessfullyCreated(user User) Message {
	return Message{"Successfully created", "User with EMAIL: " + user.Email + " created", "users"}
}

func ErrorSystemProblems(message, backTo string) Message {
	return Message{"System problems", message, backTo}
}

func ProductSuccessfullyDeleted() Message {
	return Message{"Deleted", "Product successfully deleted.", "products"}
}
