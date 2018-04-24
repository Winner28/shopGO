package service

import (
	"dao"
	"log"
	"managers"
	"net/http"
)

type Secure struct {
	roleDAO *dao.RoleDAO
	userDAO *dao.UserDAO
}

type SecureService interface {
	checkIfAdmin(r *http.Request) bool
	checkIfUser(r *http.Request) bool
}

var secureService *Secure

const ROLE_ADMIN = "admin"
const ROLE_USER = "user"

func init() {
	secureService = new(Secure)
	secureService.roleDAO = dao.GetRoleDAO()
	secureService.userDAO = dao.GetUserDAO()
}

func getSecureService() SecureService {
	return secureService
}

func GetSecureService() SecureService {
	return secureService
}

func (secure *Secure) checkIfAdmin(r *http.Request) bool {

	if !managers.GetSessionManager().UserLoggedIn(r) {
		return false
	} else {
		email := managers.GetSessionManager().GetUserEmail(r)
		role := secure.checkUserRole(email)
		log.Println(role + "asdasd")

		if role == ROLE_ADMIN {
			return true
		} else {
			return false
		}
	}
}

func (secure *Secure) checkIfUser(r *http.Request) bool {
	if !managers.GetSessionManager().UserLoggedIn(r) {
		return false
	} else {
		email := managers.GetSessionManager().GetUserEmail(r)
		role := secure.checkUserRole(email)
		if role == ROLE_USER {
			return true
		} else {
			return false
		}
	}
}

func (secure *Secure) checkUserRole(email string) string {
	user, _ := secure.userDAO.GetUserByEmail(email)
	role, _ := secure.roleDAO.Get(user.ID)
	return role.Name
}
