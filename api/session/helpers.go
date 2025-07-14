package session

import (
	"net/http"
	"time"
)

func writeCookieIfNecessary(w *sessionResponseWriter) {
	if w.done {
		return
	}

	session, ok := w.request.Context().Value(sessionContextKey{}).(*Session)
	if !ok {
		panic("session not found in request context")
	}

	cookie := &http.Cookie{
		Name:     w.sessionManager.cookieName,
		Value:    session.id,
		Domain:   w.sessionManager.domain,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(w.sessionManager.idleExpiration),
		MaxAge:   int(w.sessionManager.idleExpiration / time.Second),
	}

	http.SetCookie(w.ResponseWriter, cookie)

	w.done = true
}

func GetSession(r *http.Request) *Session {
	// .(*Session) is a type assertion, which affects `ok`
	session, ok := r.Context().Value(sessionKey).(*Session)
	if !ok {
		panic("session not found in context")
	}

	return session
}
