package pkg

import (
	"as/pkg/env"
	"as/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ERR error

type DSN struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
}

func (d *DSN) string() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", d.Host, d.User, d.Password, d.DBname, d.Port)
}

func Conn(dsn DSN) (*gorm.DB, error) {
	DB, ERR = gorm.Open(postgres.Open(dsn.string()), &gorm.Config{})
	utils.Error(ERR)

	return DB, ERR
}

func ConnQD() (*gorm.DB, error) {
	/* set up the DSN for the database */
	dsn := DSN{
		Host:     env.DBQDH,
		User:     env.DBQDU,
		Password: env.DBQDP,
		DBname:   env.DBQDD,
		Port:     env.DBQDPP,
	}

	/* connect to the database */
	DB, ERR = Conn(dsn)
	utils.Error(ERR)

	return DB, ERR
}
