package main

import (
	"database/sql"
	"log"
	"movie_premiuem/entity"
	"movie_premiuem/entity/repositories"
	"movie_premiuem/services"

	"movie_premiuem/db/migrate"

	_ "github.com/mattn/go-sqlite3"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	db, err := sql.Open("sqlite3", "./movie_premium.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ensure the database connection is valid
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	migrateUpErr := migrate.UP(db)
	if migrateUpErr != nil {
		log.Fatalf("failed to run migrations: %v", migrateUpErr)
	}

	// Initialize repository and services
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	user := entity.User{
		ID:       0,
		Email:    "aliaghdam.erfan2@gmail.com",
		Password: "Pass123!",
	}
	user, registerUserErr := userService.RegisterUser(user)
	if registerUserErr != nil {
		log.Fatalf("failed to register user: %v", registerUserErr)
	}

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
