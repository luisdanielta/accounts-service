package schemas

import (
	"as/pkg"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null"`
}

type Email struct {
	Value string `json:"value"`
}

type Password struct {
	Value string `json:"value"`
}

var db, _ = pkg.ConnQD()

type User struct {
	gorm.Model
	Id       int      `gorm:"primaryKey;autoIncrement"`
	Name     string   `gorm:"type:varchar(255);not null"`
	LastName string   `gorm:"type:varchar(255);not null"`
	Username string   `gorm:"type:varchar(255);not null;unique"`
	Email    Email    `gorm:"type:varchar(255);not null;unique"`
	Password Password `gorm:"type:varchar(255);not null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Insert() error {
	err := db.Create(&u).Error
	return err
}

/* funcs for Email */
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

func Find(item string, value string) bool {
	var user User
	err := db.Where(fmt.Sprintf("%s = ?", item), value).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

/*
func (u *User) Migrate(a bool) error {
	if a {
		db, _ := pkg.ConnQD()
		err := db.AutoMigrate(&User{})
		return err
	}
	return nil
}
*/
