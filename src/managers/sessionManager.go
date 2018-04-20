package managers

import (
	"net/http"
)

type SessionManager struct {
	name     string
	sessions map[string]bool
}

const MANAGER_NAME = "Session Manager v0.0.1"

var sessionManager *SessionManager

func init() {
	sessionManager = new(SessionManager)
	sessionManager.name = "Session Manager v0.0.1"
	sessionManager.sessions = make(map[string]bool)
}

func GetSessionManager() *SessionManager {
	return sessionManager
}

func (manager *SessionManager) CreateSession(cookie *http.Cookie) {
	manager.sessions[cookie.Value] = true
}

func (manager *SessionManager) GetUserEmailByCookie(cookie *http.Cookie) (string, bool) {
	if manager.CheckIfUserLoggedIn(cookie) {
		return cookie.Value, true
	}

	return "not logged in", false
}

func (manager *SessionManager) CheckIfUserLoggedIn(cookie *http.Cookie) bool {
	if loggedIN, contains := manager.sessions[cookie.Value]; !loggedIN || !contains {
		return false
	}
	return true
}

func (manager *SessionManager) InvalidateSession(cookie *http.Cookie) {
	manager.sessions[cookie.Value] = false
}
