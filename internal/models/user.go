package models

import "golang.org/x/crypto/bcrypt"

type Email struct {
	Value string `json:"value"`
}

func (e *Email) Validate() (bool, error) {
	/* check out Email */
	index := 0
	for i, v := range e.Value {
		if v == '@' {
			index = i
			break
		}
	}
	return e.Value[index:] == "@quantum-fsd.com", nil
}

type Password struct {
	Value string `json:"value"`
}

func (p *Password) Encrypt() (string, error) {
	/* encrypt password */
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p *Password) Decrypt(hash string) error {
	/* decrypt password */
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p.Value))
}

type User struct {
	Name     string   `json:"name"`
	LastName string   `json:"lastname"`
	Username string   `json:"username"`
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}
