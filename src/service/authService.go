package service

import (
	"log"
	"managers"
	"model"
	"net/http"
)

type authDAO interface {
	Login(email, passwordHash string) (bool, string)
	Register(model.User) (bool, string)
}

const COOKIE_NAME = "LOGGED_IN"

type AuthService struct {
	DAO authDAO
}

func GetAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (auth *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login method")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if !validateUserSignInInput(email, password) {
		log.Println("Bad user input!")
		http.Redirect(w, r, "http://localhost:8080/signin", http.StatusFound)
		return
	}
	success, message := auth.DAO.Login(email, password)
	if !success {
		http.Error(w, message, http.StatusForbidden)
	} else {
		auth.createSession(w, email)
		http.Redirect(w, r, "http://localhost:8080/profile", http.StatusFound)
	}
}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout method")
	managers.GetSessionManager().ClearSession(w)
}

func (auth *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	user := auth.parseUserFromRequest(r)
	if !validateUserSignUpInput(user) {
		log.Println("bad user input...")
		http.Redirect(w, r, "http://localhost:8080/signup", http.StatusFound)
		return
	}
	success, msg := auth.DAO.Register(user)
	if !success {
		log.Println("smth wents wrong...", msg)
	} else {
		log.Println(msg)
		auth.createSession(w, user.Email)
	}
}

func (auth *AuthService) parseUserFromRequest(r *http.Request) model.User {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	return model.User{
		FirstName:    firstname,
		LastName:     lastname,
		Email:        email,
		PasswordHash: password,
	}
}

func (auth *AuthService) createSession(w http.ResponseWriter, cookieValue string) {
	managers.GetSessionManager().SetSession(cookieValue, w)
	log.Println("Cookie is created, session is created")
}

func validateUserSignUpInput(user model.User) bool {
	return !(user.FirstName == "" || user.LastName == "" ||
		user.Email == "" || user.PasswordHash == "")
}

func validateUserSignInInput(email, password string) bool {
	return !(email == "" || password == "")
}
