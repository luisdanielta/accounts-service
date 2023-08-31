package models

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Email string
type Password string
type Username string

type User struct {
	gorm.Model
	Name     string   `gorm:"type:varchar(255);not null" json:"name"`
	LastName string   `gorm:"type:varchar(255);not null" json:"lastname"`
	Username Username `gorm:"type:varchar(255);not null" json:"username"`
	Email    Email    `gorm:"type:varchar(255);not null" json:"email"`
	Password Password `gorm:"type:varchar(255);not null" json:"password"`
}

/* funcs password type */
func (p *Password) Encrypt() (string, error) {
	/* encrypt password */
	pass := string(*p)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p *Password) Compare(plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(*p), []byte(plainPassword))
}

/* funcs email type */
func (e *Email) Validate(db *gorm.DB) (string, bool) {
	email := string(*e)

	if !CheckItemExists(db, "email", email) {
		return "Email already exists.", false
	}

	/* Validate that the email has the correct domain */
	if !strings.HasSuffix(email, "@quantum-fsd.com") {
		return "Email is not valid.", false
	}

	return "", true
}

/* funcs username type */
func (u *Username) Validate(db *gorm.DB) (string, bool) {
	/* get username */
	username := string(*u)

	/* check if the username already exists */
	if !CheckItemExists(db, "username", username) {
		return "Username already exists.", false
	}

	return "", true
}

func CheckItemExists(db *gorm.DB, field string, value string) bool {
	var user User
	r := db.Where(fmt.Sprintf("%s = ?", field), value).First(&user)
	return r.RowsAffected == 0

}
