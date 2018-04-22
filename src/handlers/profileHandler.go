package handlers

import (
	"dao"
	"net/http"
	"service"
)

var profileDAO *dao.ProfileDAO
var profileService *service.ProfileService

func init() {
	profileDAO = dao.GetProfileDAO()
	profileService = service.GetProfileService(profileDAO)
}

func (app *App) GetProfile(w http.ResponseWriter, r *http.Request) {
	profileService.GetProfile(w, r)
}

func (app *App) EditProfile(w http.ResponseWriter, r *http.Request) {

}

func (app *App) DeleteProfile(w http.ResponseWriter, r *http.Request) {

}
