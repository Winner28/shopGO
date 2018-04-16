package service

import (
	"encoding/json"
	"model"
	"net/http"
	"response"
	"strconv"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request, value interface{}) {
	params := mux.Vars(r)
	ID := params["id"]

	if user, ok := value.(model.User); ok {

	} else if product, ok := value.(model.Product); ok {

	} else {
		// ...
	}

	valueT.ID, _ = strconv.Atoi(ID)
	if err := json.NewDecoder(r.Body).Decode(&valueT); err != nil {
		response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
	}

	value, err := repository.Create(valueT)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusCreated, value)
}

func Get(w http.ResponseWriter, r *http.Request, value interface{}) {
	params := mux.Vars(r)
	ID := params["id"]
	user, err := repository.Get(ID, value)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, user)
}

func Delete(w http.ResponseWriter, r *http.Request, value interface{}) {
	params := mux.Vars(r)
	ID := params["id"]
	if err := repository.Delete(ID, value); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "User deleted")
}

func Update(w http.ResponseWriter, r *http.Request, value interface{}) {
	response.RespondJSON(w, http.StatusOK, "User deleted")
}

func GetAll(w http.ResponseWriter, r *http.Request, value interface{}) {
	var users []model.User
	if err := getConnection().DB.Find(&users).Error; err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, users)
}
