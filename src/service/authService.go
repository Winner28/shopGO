package service

import (
	"log"
	"managers"
	"model"
	"net/http"
)

type authDAO interface {
	Login(email, passwordHash string) (bool, string)
	Register(model.User) (bool, error)
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
		return
	} else {
		cookie := &http.Cookie{
			Name:  COOKIE_NAME,
			Value: email,
		}
		http.SetCookie(w, cookie)
		managers.GetSessionManager().CreateSession(cookie)
		log.Println("Cookie is created, session is created")
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
	auth.DAO.Register(model.User{})
}
