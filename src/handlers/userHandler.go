package handlers

import (
	"dao"
	"net/http"
	"service"
)

var userDAO *dao.UserDAO
var userService *service.UserService

func init() {
	userDAO = dao.GetUserDAO()
	userService = service.GetUserService(userDAO)
}

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	userService.Get(w, r)
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	userService.Create(w, r)
}

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userService.Delete(w, r)
}

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userService.Update(w, r)
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userService.GetAll(w, r)
}
