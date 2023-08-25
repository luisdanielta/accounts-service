package schemas

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null"`
}

type User struct {
	gorm.Model
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(255);not null"`
	LastName   string `gorm:"type:varchar(255);not null"`
	Username   string `gorm:"type:varchar(255);not null;unique"`
	Email      string `gorm:"type:varchar(255);not null;unique"`
	Password   string `gorm:"type:varchar(255);not null"`
	CreationAt string `gorm:"type:timestamp;not null"`
	UpdatedAt  string `gorm:"type:timestamp;not null"`
	DeletedAt  string `gorm:"type:timestamp;not null"`
}
