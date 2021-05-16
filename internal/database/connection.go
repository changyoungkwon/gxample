package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

// DBConn returns a postgres connection pool.
func DBConn() *gorm.DB {
	return conn
}

func init() {
	var err error
	conn, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal errors making connection to database %s", err))
	}
}
