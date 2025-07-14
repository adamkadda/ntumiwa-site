package session

import (
	"context"
	"log"
	"net/http"
	"time"
)

// TODO: Set up manager instance from config

// SessionManager coordinates all session operations
type SessionManager struct {
	store              SessionStore
	idleExpiration     time.Duration
	absoluteExpiration time.Duration
	cookieName         string
	domain             string
	csrfFormKey        string
	csrfHeader         string
}

func NewSessionManager(
	store SessionStore,
	gcInterval,
	idleExpiration,
	absoluteExpiration time.Duration,
	cookieName string,
	domain string,
) *SessionManager {
	m := &SessionManager{
		store:              store,
		idleExpiration:     idleExpiration,
		absoluteExpiration: absoluteExpiration,
		cookieName:         cookieName,
		domain:             domain,
	}

	go m.gc(gcInterval)

	return m
}

func (m *SessionManager) gc(d time.Duration) {
	ticker := time.NewTicker(d)

	for range ticker.C {
		m.store.gc(m.idleExpiration, m.absoluteExpiration)
	}
}

func (m *SessionManager) validate(session *Session) bool {
	if time.Since(session.createdAt) > m.absoluteExpiration ||
		time.Since(session.lastActivityAt) > m.idleExpiration {

		// Invoke the SessionStore's destroy method
		err := m.store.destroy(session.id)
		if err != nil {
			panic(err)
		}

		return false
	}

	return true
}

/*
	This is a type used to prevent key collisions in the context.

	When using context.WithValue(), it is best practice to assign
	a unique key type, to prevent accidental conflicts with keys
	used elsewhere.
*/

type sessionContextKey struct{}

/*
	By instantiating a variable of an empty struct type, we can
	create a unique context key, scoped to this package. Since it
	is an empty struct, it doesn't use memory.
*/

var sessionKey = sessionContextKey{}

/*
	We implement a start method that retrieves the session by reading
	the session cookie or generates a new one if needed.

	It then attaches the session to the request using context values.
*/

func (m *SessionManager) start(r *http.Request) (*Session, *http.Request) {
	var session *Session

	/*
		A cookie is a single key=value pair. However the SessionManager
		refers to the key in question as a cookieName, hence the field.
	*/

	// https://pkg.go.dev/net/http#Request.Cookie
	// tldr: returns the named cookie
	cookie, err := r.Cookie(m.cookieName)

	/*
		Upon succeeding, we call the Store's read() method to access
		the desired session via it's sessionID, i.e. the cookie.Value
	*/
	if err == nil {
		session, err = m.store.read(cookie.Value)
		if err != nil {
			log.Printf("Failed to read session from store: %v", err)
		}
	}

	if session == nil || !m.validate(session) {
		session = NewSession()
	}
	/*
		Contexts are Go's way to pass values, deadlines, and cancel
		signals across goroutines and handlers. They are not part of
		actual HTTP requests, they exist entirely inside the app.

		The context exists in Go memory, and are created and extended
		by Go when handling a request.

		Contexts should be only used to store request-scoped variables.
	*/

	/*
		Creates a new context, derived from an existing one.
		Essentially preparing a new context with an additional
		key=value pair, i.e. sessionKey=session
	*/
	ctx := context.WithValue(r.Context(), sessionKey, session)
	r = r.WithContext(ctx)

	return session, r
}

func (m *SessionManager) save(session *Session) error {
	session.lastActivityAt = time.Now()

	err := m.store.write(session)
	if err != nil {
		return err
	}

	return nil
}

/*
	The migrate method is called when a user's priviledge level
	changes during a session.

	The old session is deleted, and the session's ID is changed to
	ensure that any previously compromised/invalid IDs are invalidated.
*/

func (m *SessionManager) migrate(session *Session) error {
	err := m.store.destroy(session.id)
	if err != nil {
		return err
	}

	session.id = generateSessionID()

	return nil
}

/*

 */

type sessionResponseWriter struct {
	http.ResponseWriter                 // embedded to satisfy the http.Handler interface
	sessionManager      *SessionManager // used to access session/cookie logic
	request             *http.Request   // needed for context/session info
	done                bool            // tracks whether the cookie has been written
}

func (m *SessionManager) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			Start the session.

			The start method retrieves the incoming request cookie, or generates a new one.
			Very importantly, it then attaches the session to the context, which is then
			part of the new request object.
		*/
		session, request := m.start(r)

		// Create a new response writer
		sw := &sessionResponseWriter{
			ResponseWriter: w,
			sessionManager: m,
			request:        r,
		}

		/*
			Add essential headers.

			The first tells caches that the response may vary based
			on the Cookie header, so they shouldn't serve the same
			cached response to all users. Instead, prepare a separate
			cache for every variant (i.e. Cookie).

			The second prevents caching of responses that set cookies,
			forcing caches to revalidate with the origin each time
			before reusing them.
		*/
		w.Header().Add("Vary", "Cookie")
		w.Header().Add("Cache-Control", `no-cache="Set-Cookie"`)

		/*
			Before performing state-changing requests, we first
			verify the CSRF token, and fail the request if the token
			does not match what is inside the session.
		*/

		if request.Method == http.MethodPost ||
			request.Method == http.MethodPut ||
			request.Method == http.MethodPatch ||
			request.Method == http.MethodDelete {

			if !m.verifyCSRFToken(session, request) {
				http.Error(sw, "CSRF Token mismatch", http.StatusForbidden)
			}
		}

		// Call the next handler
		next.ServeHTTP(sw, request)

		// Save the session
		m.save(session)

		// Write cookie to response if needed
		writeCookieIfNecessary(sw)
	})
}

func (m *SessionManager) verifyCSRFToken(session *Session, r *http.Request) bool {
	sessionCSRFToken, ok := session.Get("csrf_token").(string)
	if !ok {
		return false
	}

	requestCSRFToken := r.FormValue(m.csrfFormKey)

	// NOTE: Just trying to keep it frontend agnostic
	if requestCSRFToken == "" {
		requestCSRFToken = r.Header.Get(m.csrfHeader)
	}

	return sessionCSRFToken == requestCSRFToken
}
