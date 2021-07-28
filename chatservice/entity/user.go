package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User data
type User struct {
	ID        ID
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
}

//NewUser creates a new User
func NewUser(email, password, name string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u User) Validate() error {
	if u.Email == "" || u.Name == "" || u.Password == "" {
		return ErrInvalidEntity
	}
	return nil
}

//ValidatePassword validate user password
func (u User) validatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
