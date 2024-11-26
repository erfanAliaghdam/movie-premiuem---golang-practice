package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"movie_premiuem/custom_errors"
	"movie_premiuem/entity"
	"movie_premiuem/entity/repositories"
	"movie_premiuem/services"
	"movie_premiuem/utils"
	"strconv"

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

	random_email_prefix := strconv.Itoa(rand.Int())
	user := entity.User{
		ID:       0,
		Email:    fmt.Sprintf("aliaghdam.erfan%s@gmail.com", random_email_prefix),
		Password: "Pass123!",
	}
	user, registerUserErr := userService.RegisterUser(user)
	if registerUserErr != nil {
		log.Fatalf("failed to register user: %v", registerUserErr)
	}

	access, refresh, jwtGenerateErr := utils.GenerateJWT(user.ID)
	if jwtGenerateErr != nil {
		log.Fatalf("failed to generate JWT: %v", jwtGenerateErr)
	}
	fmt.Printf("user refresh %v \n user access %v \n", refresh, access)

	verified, verifyJwtErr := utils.VerifyToken(access)
	if verifyJwtErr != nil {
		log.Fatalf("failed to verify JWT: %v", verifyJwtErr)
	}
	fmt.Printf("user verified: %v \n", verified)

	// test order
	fmt.Println("----------")
	// initialize order repo and service
	order := entity.Order{
		ID:        0,
		UserID:    user.ID,
		Paid:      false,
		PaidPrice: 500.0,
	}
	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	order, orderCreateErr := orderService.CreateOrder(order)
	if orderCreateErr != nil {
		log.Fatalf("failed to create order: %v", orderCreateErr)
	}
	fmt.Printf("order created: %v", order)

	//test user license
	fmt.Println("----------")
	license := entity.License{
		ID:          0,
		Title:       "license regular",
		FinishMonth: 2,
		Price:       2,
		LicenseType: entity.RegularLicenseType,
	}

	licenseRepository := repositories.NewLicenseRepository(db)
	license, createLicenseErr := licenseRepository.CreateLicense(license)
	if createLicenseErr != nil {
		log.Fatalf("failed to create license: %v", createLicenseErr)
	}
	fmt.Printf("license created: %v\n", license)

	fmt.Println("----------")
	//test user license
	userLicenseRepository := repositories.NewUserLicenseRepository(db)
	userLicenseService := services.NewUserLicenseService(userLicenseRepository)
	userLicense, userLicenseCreateErr := userLicenseService.CreateUserLicense(license.ID, user.ID)
	if userLicenseCreateErr != nil {
		if errors.Is(userLicenseCreateErr, custom_errors.ErrUserHasActiveLicense) {
			log.Fatalf("user license already exists: %v", userLicenseCreateErr)
		}
		log.Fatalf("failed to create user license: %v", userLicenseCreateErr)
	}

	fmt.Printf("license created for user : %v\n", userLicense)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
