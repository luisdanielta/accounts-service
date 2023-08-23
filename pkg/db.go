package pkg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ( // dsn is a string that contains the database connection information
	dsn string = "host=localhost user=postgres password=postgres dbname=csh_users port=5432 sslmode=disable"
	db  *gorm.DB
	err error
)

func Conn() (*gorm.DB, error) { // Conn is a function that returns a pointer to a gorm.DB and an error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
