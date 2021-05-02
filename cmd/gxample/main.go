package main

import (
	"fmt"

	"github.com/changyoungkwon/gxample/internal/config"
	"github.com/changyoungkwon/gxample/internal/database"
)

func main() {
	conn := database.DBConn()
	if err := database.CheckConn(conn); err != nil {
		fmt.Printf("hello world")
	}
	fmt.Println(config.Get().Database.URL)
}
