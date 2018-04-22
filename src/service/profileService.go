package service

import (
	"log"
	"managers"
	"model"
	"net/http"
	"resources"
)

type profileDAO interface {
	GetProfile(email string) (model.User, error)
	UpdateProfile(user model.User) (model.User, error)
	DeleteProfile(user model.User) error
}

type ProfileService struct {
	DAO profileDAO
}

func GetProfileService(dao profileDAO) *ProfileService {
	return &ProfileService{dao}
}

func (service *ProfileService) GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Get profile")
	if cookie, _ := r.Cookie(COOKIE_NAME); cookie != nil {
		if email, IN := managers.GetSessionManager().GetUserEmailByCookie(cookie); IN {
			log.Println("LOGGED IN!")

			profile, err := service.DAO.GetProfile(email)
			if err != nil {
				if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetInternalServerErrorError()); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
			if err := resources.GetTemplatesContainer().GetTemplate("profile").Execute(w, profile); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	log.Println("Cant get profile, login first")

	if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetError(http.StatusForbidden, "You need to login first!")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
