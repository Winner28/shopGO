package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var app App

func Init() {
	app.Router = mux.NewRouter()
	app.setRouters()
	http.ListenAndServe(":8080", app.Router)
}

func (app *App) setRouters() {

	// users routes
	app.Router.HandleFunc("/users/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/users/{id}", app.CreateUser).Methods("POST")
	app.Router.HandleFunc("/users/{id}", app.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id}", app.DeleteUser).Methods("DELETE")
	app.Router.HandleFunc("/users", app.GetAllUsers).Methods("GET")

}
