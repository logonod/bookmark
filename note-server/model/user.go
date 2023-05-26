package model

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
}

func ComparePasswordHash(hashedPassword, givenPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, givenPassword)
	return err == nil
}

type User struct {
	Model `bson:",inline"`

	Phone          string `json:"phone,omitempty"`
	HashedPassword string `json:"password,omitempty" bson:"password,omitempty"`
}

func (u *User) SetPassword(password string) error {
	hashed, err := GeneratePasswordHash([]byte(password))
	if err != nil {
		return err
	}
	u.HashedPassword = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return ComparePasswordHash([]byte(u.HashedPassword), []byte(password))
}
