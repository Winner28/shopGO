package service

import (
	"encoding/json"
	"model"
	"net/http"
	"response"
	"strconv"

	"github.com/gorilla/mux"
)

type userDAO interface {
	Get(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(ID int, user model.User) (model.User, error)
	Delete(ID int) error
	FindAll() ([]model.User, error)
}

type UserService struct {
	DAO userDAO
}

func GetUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

func (service *UserService) Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
		return
	}
	user.ID = ID
	user, err := service.DAO.Create(user)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusCreated, user)
}

func (service *UserService) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	user, err := service.DAO.Get(ID)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, user)
}

func (service *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	if err := service.DAO.Delete(ID); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "User Successfully Deleted")
}

func (service *UserService) Update(w http.ResponseWriter, r *http.Request) {
	response.RespondJSON(w, http.StatusOK, "Value Updated")
}

func (service *UserService) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := service.DAO.FindAll()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, users)
}
