package service

import (
	"encoding/json"
	"model"
	"net/http"
	"response"

	"github.com/gorilla/mux"
)

type UserRepository interface {
	Create(user model.User)
	Get(ID int)
	Delete(ID int)
	Update(ID int, user model.User)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
	}
	if err = getConnection().DB.Create(&user).Error; err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	user := model.User{}
	if err := getConnection().DB.First(&user, ID).Error; err != nil {
		response.RespondError(w, http.StatusNotFound, "User not found")
	}
	response.RespondJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	user := model.User{}
	if err := getConnection().DB.Delete(user, ID).Error; err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	if err := getConnection().DB.Find(&users).Error; err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, users)
}
