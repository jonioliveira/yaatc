package user

import (
	"time"

	"github.com/jonioliveira/yaatc/internal/id"
	"golang.org/x/crypto/bcrypt"
)

type Builder interface {
	AddId() Builder
	SetEmail(email string) Builder
	SetPassword(password string) Builder
	SetFirstName(firstName string) Builder
	SetLastName(lastName string) Builder
	Build() User
}

type UserBuilder struct {
	User
}

func NewBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) AddId() *UserBuilder {
	b.ID = id.NewID()
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	password, err := generatePassword(password)
	if err != nil {
		b.Password = ""
	}
	b.Password = password
	return b
}

func (b *UserBuilder) SetFirstName(firstName string) *UserBuilder {
	b.FirstName = firstName
	return b
}

func (b *UserBuilder) SetLastName(lastName string) *UserBuilder {
	b.LastName = lastName
	return b
}

func (b *UserBuilder) Build() *User {
	return &User{
		ID:        b.ID,
		Email:     b.Email,
		Password:  b.Password,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: b.UpdatedAt,
	}
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
