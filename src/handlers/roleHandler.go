package handlers

import (
	"dao"
	"net/http"
	"service"
)

var roleDAO *dao.RoleDAO
var roleService *service.RoleService

func init() {
	roleDAO = dao.GetRoleDAO()
	roleService = service.GetRoleService(roleDAO)
}

func (app *App) GetUserRole(w http.ResponseWriter, r *http.Request) {
	roleService.Get(w, r)
}

func (app *App) CreateRole(w http.ResponseWriter, r *http.Request) {
	roleService.Create(w, r)
}

func (app *App) DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	roleService.Delete(w, r)
}

func (app *App) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	roleService.Update(w, r)
}

func (app *App) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roleService.GetAll(w, r)
}
