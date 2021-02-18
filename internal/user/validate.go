package user

import (
	"golang.org/x/crypto/bcrypt"

	. "github.com/jonioliveira/yaatc/internal/error"
)

func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

func (u *User) ValidatePassword() error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
