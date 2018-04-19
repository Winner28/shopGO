package managers

import "net/http"

type SessionManager struct {
	name     string
	sessions map[string]bool
}

const COOKIE_NAME = "LOGGED_IN"
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
	if cookie.Name == COOKIE_NAME {
		manager.sessions[cookie.Value] = true
	}
}

func (manager *SessionManager) GetUserEmailByCookie(cookie *http.Cookie) (string, bool) {
	if cookie.Name == COOKIE_NAME {
		loggedIN, contains := manager.sessions[cookie.Value]
		if !contains && !loggedIN {
			return "#empty", false
		}
		return cookie.Value, true
	}

	return "#empty", false
}

func (manager *SessionManager) DeleteSession(cookie *http.Cookie) {
	manager.sessions[cookie.Value] = false
}
