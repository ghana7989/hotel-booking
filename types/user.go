package types

import (
	"net/mail"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	minNameLength     = 3
	minPasswordLength = 6
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	// encrypt the password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	params.Password = string(encryptedPassword)
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}

func (params *CreateUserParams) Validate() error {
	if len(params.FirstName) < minNameLength {
		return errors.New("first name must be at least 3 characters")
	}
	if len(params.LastName) < minNameLength {
		return errors.New("last name must be at least 3 characters")
	}
	if len(params.Password) < minPasswordLength {
		return errors.New("password must be at least 6 characters")
	}
	if _, err := mail.ParseAddress(params.Email); err != nil {
		return errors.New("invalid email address")
	}
	return nil
}
