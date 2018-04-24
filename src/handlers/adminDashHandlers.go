package handlers

import "net/http"

func (app *App) UpdateUserForm(w http.ResponseWriter, r *http.Request) {
	userService.UpdateForm(w, r)
}
