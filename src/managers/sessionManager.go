package managers

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

type SessionManager struct {
	name string
}

const MANAGER_NAME = "Session Manager v0.0.1"

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var sessionManager *SessionManager

func init() {
	sessionManager = new(SessionManager)
	sessionManager.name = "Session Manager v0.0.1"
}

func GetSessionManager() *SessionManager {
	return sessionManager
}

func (manager *SessionManager) SetSession(userEmail string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userEmail,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func (manager *SessionManager) ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func (manager *SessionManager) GetUserEmail(request *http.Request) (userEmail string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userEmail = cookieValue["name"]
		}
	}
	return userEmail
}

// Returns true if User Logged In
func (manager *SessionManager) UserLoggedIn(request *http.Request) bool {
	return manager.GetUserEmail(request) != ""
}
