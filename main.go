package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/routes"
	"net/http"
	"time"
)

func main() {
	core.InitApplication("./movie_premium.db", "localhost:6379")
	defer core.AppInstance.CloseDB()

	db, err := sql.Open("sqlite3", "./movie_premium.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
