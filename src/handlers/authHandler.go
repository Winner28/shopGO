package handlers

import (
	"dao"
	"log"
	"managers"
	"net/http"
	"resources"
	"service"
)

var authDAO *dao.AuthDAO
var authService *service.AuthService

const COOKIE_NAME = "LOGGED_IN"

func init() {
	authDAO = dao.GetAuthDAO()
	authService = service.GetAuthService(authDAO)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	authService.Login(w, r)
}

func (app *App) LoginPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Login page")
	if cookie, _ := r.Cookie(COOKIE_NAME); cookie != nil {
		if managers.GetSessionManager().CheckIfUserLoggedIn(cookie) {
			log.Println("User already logged in")
			return
		}
	}
	if err := resources.GetTemplatesContainer().GetTemplate("signin").Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	authService.Logout(w, r)
}

func (app *App) SignUp(w http.ResponseWriter, r *http.Request) {
	authService.Register(w, r)
}

func (app *App) SignUpPage(w http.ResponseWriter, r *http.Request) {
	log.Println("SignUp page")
	if cookie, _ := r.Cookie(COOKIE_NAME); cookie != nil {
		if managers.GetSessionManager().CheckIfUserLoggedIn(cookie) {
			log.Println("User already logged in, redirecting to home page.")
			//	http.Redirect(w, r, "http://localhost:8080/profile", http.StatusFound)
			return
		}
	}
	if err := resources.GetTemplatesContainer().GetTemplate("signup").Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
