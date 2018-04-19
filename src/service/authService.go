package service

import (
	"log"
	"managers"
	"model"
	"net/http"

	"github.com/gorilla/sessions"
)

type authDAO interface {
	Login(email, passwordHash string) (bool, string)
	Register(model.User) (bool, error)
}

var sessionStore = sessions.NewCookieStore([]byte("#Encrypted"))

const COOKIE_NAME = "LOGGED_IN"

type AuthService struct {
	DAO authDAO
}

func GetAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (auth *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	email := r.FormValue("email")
	password := r.FormValue("password")
	if cookie == nil {
		log.Println("nil cookie!!!")
		signIn, msg := auth.DAO.Login(email, password)
		if !signIn {
			log.Println("problemos", msg)
			http.Redirect(w, r, "http://localhost:8080/signin", 301)
		} else {
			cookie = &http.Cookie{
				Name:  COOKIE_NAME,
				Value: email,
			}
			http.SetCookie(w, cookie)
			managers.GetSessionManager().CreateSession(cookie)
		}
	} else {
		if _, loggedIn := managers.GetSessionManager().GetUserEmailByCookie(cookie); !loggedIn {
			log.Println("nononononon")
			http.Redirect(w, r, "http://localhost:8080/signin", 301)
			//TODO AUTH ON LOG
		}
		log.Println("you already signed in")
		http.Redirect(w, r, "http://localhost:8080/users", 301)
	}
}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	if cookie, _ := r.Cookie(COOKIE_NAME); cookie != nil {
		managers.GetSessionManager().DeleteSession(cookie)
		log.Println("Log out!")
		http.Redirect(w, r, "http://localhost:8080/signin", 301)
	}

	log.Println("You need sign in first => then logout!")
	http.Redirect(w, r, "http://localhost:8080/signin", 301)
}

func (auth *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	auth.DAO.Register(model.User{})
}
