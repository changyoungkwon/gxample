package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg"
)

var conn *pg.DB

// DBConn returns a postgres connection pool.
func DBConn() *pg.DB {
	return conn
}

// CheckConn check if connections vlaid
func CheckConn(db *pg.DB) error {
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	return err
}

func init() {
	opt, err := pg.ParseURL(databaseURL)
	if err != nil {
		panic(fmt.Errorf("fatal errors making connections to database %s", err))
	}

	conn = pg.Connect(opt)
	if err := CheckConn(conn); err != nil {
		panic(fmt.Errorf("fatal errors making connection to database %s", err))
	}
}
