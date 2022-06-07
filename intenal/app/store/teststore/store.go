package teststore

import (
	"github.com/honyshyota/http-rest-api/intenal/app/model"
	"github.com/honyshyota/http-rest-api/intenal/app/store"
)

// Store ...
type Store struct {
	UserRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.UserRepository
}
