package models

import (
	"FS01/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	LastLogin string `json:"last_login" db:"last_login"`
}

// Create will create user in the db
func (u *User) CreateUser() error {
	_, err := database.DBClient.Exec("INSERT INTO users (ID, Name, Email, Password) VALUES (uuid_generate_v4(), $1, $2, $3);", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

// Hash will hash the password
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPass will check user password
func (u *User) CheckPassword(providedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}
