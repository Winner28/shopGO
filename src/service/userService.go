package service

import (
	"encoding/json"
	"log"
	"model"
	"net/http"
	"resources"
	"response"
	"strconv"

	"github.com/gorilla/mux"
)

type userDAO interface {
	Get(ID int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User, role model.Role) (model.User, error)
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
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}
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
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	user, err := service.DAO.Get(ID)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := resources.GetTemplatesContainer().GetTemplate("profile").Execute(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	if err := service.DAO.Delete(ID); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}

	if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.UserSuccessfullyDeleted()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *UserService) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update method!")
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	user, role := service.getUserAndRoleFromRequest(r)
	user.ID = ID
	role.UserID = ID
	user, err := service.DAO.Update(user, role)
	if err != nil {
		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileUpdatingUser()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.UserSuccessfullyUpdated(user)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (service *UserService) getUserAndRoleFromRequest(r *http.Request) (model.User, model.Role) {
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	password := r.FormValue("password")
	email := r.FormValue("email")
	role := r.FormValue("role")
	return model.User{
			FirstName:    firstName,
			LastName:     lastName,
			Email:        email,
			PasswordHash: password,
		},
		model.Role{
			Name: role,
		}

}

func (service *UserService) UpdateForm(w http.ResponseWriter, r *http.Request) {
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])

	if user, err := service.DAO.Get(ID); err == nil {
		if err := resources.GetTemplatesContainer().GetTemplate("updateUser").Execute(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.WeDontHaveSuchUser()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func (service *UserService) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all Users method")
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
	users, err := service.DAO.FindAll()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	if err := resources.GetTemplatesContainer().GetTemplate("getAllUsers").Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
