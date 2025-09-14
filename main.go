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
	config.DB.AutoMigrate(&models.User{})
	utils.InitRedis()
	r := mux.NewRouter()

	router.UserRoutes(r, config.DB)

	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		gorillaHandlers.AllowCredentials(),
	)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(r)))
}
