package application

import "github.com/sjaureguio/eshop/internal/user/domain"

// WebHandler is a ports in
type WebHandler interface {
	Create(m *domain.User) error
	GetByEmail(email string) (domain.User, error)
	GetAll() (domain.Users, error)
}

// Storage is a ports out
type Storage interface {
	Create(m *domain.User) error
	GetByEmail(email string) (domain.User, error)
	GetAll() (domain.Users, error)
}
