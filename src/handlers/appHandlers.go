package handlers

import (
	"html/template"
	"net/http"
	"resources"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var app App
var templates *resources.Templates

func Init() {
	app.Router = mux.NewRouter()
	app.setRouters()
	app.setTemplates()
	http.ListenAndServe(":8080", app.Router)
}

func (app *App) setRouters() {
	app.setAuthRoutes()
	app.setUsersRoutes()
	app.setProductRoutes()
	app.setRoleRoutes()
	app.setCustomRoutes()
}

func (app *App) setAuthRoutes() {
	app.Router.HandleFunc("/signin", app.Login).Methods("POST")
	app.Router.HandleFunc("/signin", app.LoginPage).Methods("GET")
	app.Router.HandleFunc("/logout", app.Logout).Methods("GET")
	app.Router.HandleFunc("/signup", app.Register).Methods("POST")
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

func (app *App) setCustomRoutes() {
}

func (app *App) setTemplates() {
	templates = resources.GetTemplatesContainer()
	templates.AddTemplate("signin", template.Must(template.ParseFiles("templates/signin.html")))
	app.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
}
