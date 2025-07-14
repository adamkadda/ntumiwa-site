package session

import "time"

type SessionStore interface {
	read(id string) (*Session, error)
	write(session *Session) error // writes a session to the storage engine
	destroy(id string) error
	gc(idleExpiration, absoluteExpiration time.Duration) error
}
