package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
)

/*
	Many thanks to Mohamed Said at themsaid.com
	Would not have figured this out myself!

	https://themsaid.com/building-secure-session-manager-in-go
*/

type Session struct {
	createdAt      time.Time
	lastActivityAt time.Time
	id             string
	data           map[string]any
}

func generateSessionID() string {
	id := make([]byte, 32) // 32 * 8 = 256 bits

	_, err := io.ReadFull(rand.Reader, id)
	if err != nil {
		panic("somehow failed to generate session id")
	}

	return base64.RawURLEncoding.EncodeToString(id)
}

func generateCSRFToken() string {
	id := make([]byte, 32)

	_, err := io.ReadFull(rand.Reader, id)
	if err != nil {
		panic("somehow failed to generate CSRF token")
	}

	return base64.RawURLEncoding.EncodeToString(id)
}

func NewSession() *Session {
	return &Session{
		id:             generateSessionID(),
		data:           map[string]any{"csrf_token": generateCSRFToken()},
		createdAt:      time.Now(),
		lastActivityAt: time.Now(),
	}
}

func (s *Session) Get(key string) any {
	return s.data[key]
}

func (s *Session) Put(key string, value any) {
	s.data[key] = value
}

func (s *Session) Delete(key string) {
	delete(s.data, key)
}
