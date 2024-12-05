package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/db/migrate"
	"movie_premiuem/core/routes"
	"net/http"
	"time"
)

func main() {
	config := core.LoadConfig()

	core.InitApplication(config.DBName, config.RedisAddr)
	defer core.AppInstance.CloseDB()

	migrate.UP(core.AppInstance.GetDB())

	// run server
	srv := &http.Server{
		Addr:              ":8001",
		Handler:           routes.Routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	fmt.Println("Starting web app on port :8001")

	serverErr := srv.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
