package main

import (
	"fmt"
	"log"
	"net/http"
	"rentroom/config"
	"rentroom/models"
	"rentroom/router"
	"rentroom/utils"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.ConnectDatabase()
	err := config.DB.Exec("PRAGMA foreign_keys = ON").Error
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}
	config.DB.AutoMigrate(
		&models.User{},
		&models.Bank{},
		&models.Country{},
		&models.Amenity{},
		&models.Property{},
		&models.UserProperties{},
		&models.PropertyAmenities{},
		&models.Transaction{},
		&models.Voucher{},
		&models.Review{},
	)
	utils.SeedInitialData(config.DB)
	utils.InitRedis()
	r := mux.NewRouter()

	router.RegisterUserRoutes(r, config.DB)
	router.RegisterAdminRoutes(r, config.DB)
	router.RegisterVoucherRoutes(r, config.DB)
	router.RegisterPropertyRoutes(r, config.DB)
	router.RegisterTransactionRoutes(r, config.DB)

	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		gorillaHandlers.AllowCredentials(),
	)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(r)))
}
