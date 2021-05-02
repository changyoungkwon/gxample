package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/changyoungkwon/gxample/internal/config"
	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/routes"
)

func main() {
	conn := database.DBConn()
	if err := database.CheckConn(conn); err != nil {
		fmt.Printf("fatal error while database connection check, %s", err)
	}

	svr := &http.Server{
		Handler:      routes.Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%d", config.Get().API.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := svr.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("something went wrong, %s", err)
		}
	}
}
