package application

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sjaureguio/eshop/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) *UseCase {
	return &UseCase{storage: s}
}

func (u UseCase) Create(m *domain.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID

	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	m.Password = ""

	return nil
}

func (u UseCase) GetByEmail(email string) (domain.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s %w", "storage.GetByEmail()", err)
	}

	return user, nil
}

func (u UseCase) GetAll() (domain.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return domain.Users{}, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return users, nil
}
