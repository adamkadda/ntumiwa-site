package session

import (
	"fmt"
	"sync"
	"time"
)

/*
	This is a (hopefully) temporary in-memory session store.

	It's not persistent, which is not ideal, but it should get
	the job done while I build the more essential features.
*/

type InMemorySessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]*Session),
	}
}

func (s *InMemorySessionStore) read(id string) (*Session, error) {
	s.mu.RLock()
	defer s.mu.Unlock()

	session, ok := s.sessions[id]
	if !ok {
		return nil, fmt.Errorf("could not find session: %s", id)
	}

	return session, nil
}

func (s *InMemorySessionStore) write(session *Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.id] = session

	return nil
}

func (s *InMemorySessionStore) destroy(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.sessions[id]
	if !ok {
		return fmt.Errorf("cannot delete nonexistent session: %s", id)
	}

	delete(s.sessions, id)

	return nil
}

func (s *InMemorySessionStore) gc(tti, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, session := range s.sessions {
		if time.Since(session.lastActivityAt) > tti ||
			time.Since(session.createdAt) > ttl {
			delete(s.sessions, id)
		}
	}

	return nil
}
