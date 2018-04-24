package handlers

import "net/http"

func (app *App) UpdateUserForm(w http.ResponseWriter, r *http.Request) {
	userService.UpdateForm(w, r)
}

func (app *App) CreateUserForm(w http.ResponseWriter, r *http.Request) {
	userService.CreateForm(w, r)
}
