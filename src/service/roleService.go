package service

import (
	"encoding/json"
	"model"
	"net/http"
	"response"
	"strconv"

	"github.com/gorilla/mux"
)

type roleDAO interface {
	Get(ID int) (model.Role, error)
	Create(role model.Role) (model.Role, error)
	Update(ID int, role model.Role) (model.Role, error)
	Delete(ID int) error
	FindAll() ([]model.Role, error)
}

type RoleService struct {
	DAO roleDAO
}

func GetRoleService(dao roleDAO) *RoleService {
	return &RoleService{dao}
}

func (service *RoleService) Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])
	var role model.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
		return
	}
	role.UserID = userID
	role, err := service.DAO.Create(role)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusCreated, role)
}

func (service *RoleService) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])
	role, err := service.DAO.Get(userID)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, role)
}

func (service *RoleService) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])
	if err := service.DAO.Delete(userID); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "User Role Successfully Deleted")
}

func (service *RoleService) Update(w http.ResponseWriter, r *http.Request) {
	response.RespondJSON(w, http.StatusOK, "Value Updated")
}

func (service *RoleService) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := service.DAO.FindAll()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, roles)
}
