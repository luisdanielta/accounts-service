package pkg

import (
	"as/pkg/env"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c *Conn) GetConn() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", c.Host, c.User, c.Password, c.Database, c.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

/* database - qfsd */
var ConnDB = Conn{
	Host:     env.DBQDH,
	Port:     env.DBQDPP,
	User:     env.DBQDU,
	Password: env.DBQDP,
	Database: env.DBQDD,
}
