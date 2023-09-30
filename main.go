package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/matt9mg/rawflix-api/cmd"
	"github.com/matt9mg/rawflix-api/controllers"
	"github.com/matt9mg/rawflix-api/db"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/validators"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.NewStore(&db.StoreConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}).Connect()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewUser(conn)

	pwCost, err := strconv.Atoi(os.Getenv("PASSWORD_COST"))

	if err != nil {
		log.Fatal(err)
	}

	passwordHasher := services.NewPassword(&services.PasswordConfig{Cost: pwCost})

	if os.Getenv("CLI") == "true" {
		cmd.Execute(userRepo, passwordHasher)
		os.Exit(0)
	}

	jwt := services.NewJWT([]byte(os.Getenv("JWT_SECRET")), userRepo)

	registerValidator := validators.NewRegister(userRepo)
	loginValidator := validators.NewLogin(userRepo, passwordHasher)

	registerController := controllers.NewRegister(registerValidator, passwordHasher, userRepo)
	loginController := controllers.NewLogin(loginValidator, jwt, userRepo)
	logoutController := controllers.NewLogout(jwt, userRepo)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/register-field-data", registerController.GetRegisterFieldData)
	app.Post("/register", registerController.Register)
	app.Post("/login", loginController.Login)
	app.Post("/logout", logoutController.Logout)

	log.Fatal(app.Listen(":3002"))
}
