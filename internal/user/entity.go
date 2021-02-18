package user

import (
	"time"

	. "github.com/jonioliveira/yaatc/internal/error"
	. "github.com/jonioliveira/yaatc/internal/id"
)

type User struct {
	ID        ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []ID
}

func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := NewBuilder().AddId().SetEmail(email).SetPassword(password).SetFirstName(firstName).SetLastName(lastName).Build()
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}
