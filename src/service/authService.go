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
	success, message := auth.DAO.Login(email, password)
	if !success {
		http.Error(w, message, http.StatusForbidden)

	} else {
		auth.createSession(w, email)
		http.Redirect(w, r, "http://localhost:8080/", 301)
	}
}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout method")
	if cookie, _ := r.Cookie(COOKIE_NAME); cookie != nil {
		if loggedIN := managers.GetSessionManager().CheckIfUserLoggedIn(cookie); loggedIN {
			log.Println("logouted")
			managers.GetSessionManager().InvalidateSession(cookie)
			//http.Redirect(w, r, "http://localhost:8080/", 301)
		} else {
			log.Println("You need to login first")
			//http.Redirect(w, r, "http://localhost:8080/signin", 301)
		}
		return
	}
}

func (auth *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	success, msg := auth.DAO.Register(auth.parseUserFromRequest(r))
	if !success {
		log.Println("smth wents wrong...", msg)
		return
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
	cookie := &http.Cookie{
		Name:  COOKIE_NAME,
		Value: cookieValue,
	}
	http.SetCookie(w, cookie)
	managers.GetSessionManager().CreateSession(cookie)
	log.Println("Cookie is created, session is created")
}
