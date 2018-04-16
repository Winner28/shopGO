package handlers

import (
	"model"
	"net/http"
	"service"
)

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	service.Get(w, r, model.User{})
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	service.Create(w, r, model.User{})
}

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	service.Delete(w, r, model.User{})
}

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	service.Update(w, r, model.User{})
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	service.GetAll(w, r, model.User{})
}
