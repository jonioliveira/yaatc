package user

import (
	. "github.com/jonioliveira/yaatc/internal/id"
)

// Reader interface
type Reader interface {
	Get(id ID) (*User, error)
	Search(query string) ([]*User, error)
	List() ([]*User, error)
}

// Writer user writer
type Writer interface {
	Create(e *User) (ID, error)
	Update(e *User) error
	Delete(id ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUser(id ID) (*User, error)
	SearchUsers(query string) ([]*User, error)
	ListUsers() ([]*User, error)
	CreateUser(email, password, firstName, lastName string) (ID, error)
	UpdateUser(e *User) error
	DeleteUser(id ID) error
}

type Service struct {
	repository Repository
}

// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

// CreateUser Create an user
func (s *Service) CreateUser(email, password, firstName, lastName string) (ID, error) {
	e, err := NewUser(email, password, firstName, lastName)
	if err != nil {
		return e.ID, err
	}
	return s.repository.Create(e)
}
