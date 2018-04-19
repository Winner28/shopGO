package handlers

import (
	"dao"
	"log"
	"net/http"
	"resources"
	"service"
)

var authDAO *dao.AuthDAO
var authService *service.AuthService

func init() {
	authDAO = dao.GetAuthDAO()
	authService = service.GetAuthService(authDAO)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, "user-session")
	session.Values["auth"] = true
	session.Values["id"] = 1
	session.Save(r, w)
	authService.Login(w, r)
}

func (app *App) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := resources.GetTemplatesContainer().GetTemplate("signin").Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, "user-session")
	session.Values["auth"] = false
	session.Values["id"] = 0
	session.Save(r, w)
	authService.Logout(w, r)

}

func (app *App) Register(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, "user-session")

	// Check if user is authenticated
	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	log.Println(w, "The cake is a lie!")
	authService.Register(w, r)
}
