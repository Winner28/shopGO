package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type App struct {
	Router *mux.Router
}

var app App
var SessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

func Init() {
	app.Router = mux.NewRouter()
	app.setRouters()

	http.ListenAndServe(":8080", app.Router)
}

func (app *App) setRouters() {
	app.setAuthRoutes()
	app.setUsersRoutes()
	app.setProductRoutes()
	app.setRoleRoutes()
}

func (app *App) setAuthRoutes() {
	app.Router.HandleFunc("/login", app.Login).Methods("POST")
	app.Router.HandleFunc("/logout", app.Logout).Methods("POST")
	app.Router.HandleFunc("/register", app.Register).Methods("POST")
}

func (app *App) setUsersRoutes() {
	app.Router.HandleFunc("/users/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/users/{id}", app.CreateUser).Methods("POST")
	app.Router.HandleFunc("/users/{id}", app.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id}", app.DeleteUser).Methods("DELETE")
	app.Router.HandleFunc("/users", app.GetAllUsers).Methods("GET")
}

func (app *App) setProductRoutes() {
	app.Router.HandleFunc("/products/{id}", app.GetProduct).Methods("GET")
	app.Router.HandleFunc("/products/{id}", app.CreateProduct).Methods("POST")
	app.Router.HandleFunc("/products/{id}", app.UpdateProduct).Methods("PUT")
	app.Router.HandleFunc("/products/{id}", app.DeleteProduct).Methods("DELETE")
	app.Router.HandleFunc("/products", app.GetAllProducts).Methods("GET")
}

func (app *App) setRoleRoutes() {
	app.Router.HandleFunc("/roles/{id}", app.GetUserRole).Methods("GET")
	app.Router.HandleFunc("/roles/{id}", app.CreateRole).Methods("POST")
	app.Router.HandleFunc("/roles/{id}", app.UpdateUserRole).Methods("PUT")
	app.Router.HandleFunc("/roles/{id}", app.DeleteUserRole).Methods("DELETE")
	app.Router.HandleFunc("/roles", app.GetAllRoles).Methods("GET")
}
