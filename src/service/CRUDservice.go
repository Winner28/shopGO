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
	ID, _ := strconv.Atoi(params["id"])

	var valueType interface{}

	if user, ok := value.(model.User); ok {
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
		}
		valueType = user
	} else if product, ok := value.(model.Product); ok {
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
		}
		valueType = product
	} else {
		// ...
	}

	value, err := repository.Create(ID, valueType)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusCreated, value)
}

func Get(w http.ResponseWriter, r *http.Request, value interface{}) {
	params := mux.Vars(r)
	ID := params["id"]
	value, err := repository.Get(ID, value)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, value)
}

func Delete(w http.ResponseWriter, r *http.Request, value interface{}) {
	params := mux.Vars(r)
	ID := params["id"]
	if err := repository.Delete(ID, value); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "Value Deleted")
}

func Update(w http.ResponseWriter, r *http.Request, value interface{}) {
	response.RespondJSON(w, http.StatusOK, "Value Deleted")
}

func GetAll(w http.ResponseWriter, r *http.Request, value interface{}) {
	values, err := repository.FindAll(value)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, values)
}
