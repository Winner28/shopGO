package handlers

import "net/http"
import "service"

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	service.GetUser(w, r)
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	service.CreateUser(w, r)
}

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	service.DeleteUser(w, r)
}

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	service.UpdateUser(w, r)
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	service.GetAllUsers(w, r)
}
