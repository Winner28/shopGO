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
	if !managers.GetSessionManager().UserLoggedIn(r) {
		log.Println("User not logged in!")
		if err := resources.GetTemplatesContainer().GetTemplate("signin").Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("User already logged in...")
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

	if err := resources.GetTemplatesContainer().GetTemplate("signup").Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
