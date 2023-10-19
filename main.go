package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/matt9mg/rawflix-api/cmd"
	"github.com/matt9mg/rawflix-api/controllers"
	"github.com/matt9mg/rawflix-api/db"
	"github.com/matt9mg/rawflix-api/middleware"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/transformers"
	"github.com/matt9mg/rawflix-api/validators"
	"log"
	"net/http"
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
	movieRepo := repositories.NewMovie(conn)
	interactionRepo := repositories.NewInteraction(conn)

	pwCost, err := strconv.Atoi(os.Getenv("PASSWORD_COST"))

	if err != nil {
		log.Fatal(err)
	}

	passwordHasher := services.NewPassword(&services.PasswordConfig{Cost: pwCost})

	recombee := services.NewRecoombe(&http.Client{})

	if os.Getenv("CLI") == "true" {
		cmd.Execute(userRepo, passwordHasher, movieRepo, recombee, interactionRepo)
		os.Exit(0)
	}

	jwt := services.NewJWT([]byte(os.Getenv("JWT_SECRET")), userRepo)

	registerValidator := validators.NewRegister(userRepo)
	loginValidator := validators.NewLogin(userRepo, passwordHasher)
	interactionValidator := validators.NewInteraction(movieRepo)

	movieTransformer := transformers.NewRecommendationMovie()
	interactionTransformer := transformers.NewInteraction()

	jwtMiddleware := middleware.NewJWT(jwt)

	registerController := controllers.NewRegister(registerValidator, passwordHasher, userRepo)
	loginController := controllers.NewLogin(loginValidator, jwt, userRepo)
	logoutController := controllers.NewLogout(jwt, userRepo)
	userItemsController := controllers.NewUserItems(recombee, movieRepo, movieTransformer)
	favouriteController := controllers.NewBookmark(interactionRepo, interactionValidator, interactionTransformer, recombee)
	viewedController := controllers.NewViewed(interactionRepo, interactionValidator, interactionTransformer, recombee)
	movieController := controllers.NewMovie(movieRepo, movieTransformer)
	itemsToItemController := controllers.NewItemItems(recombee, movieRepo, movieTransformer)
	segmentsController := controllers.NewSegments(recombee)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/recommend-for-scenario", jwtMiddleware.Validate, userItemsController.RecommendForScenario)
	app.Get("/register-field-data", registerController.GetRegisterFieldData)
	app.Post("/register", registerController.Register)
	app.Post("/login", loginController.Login)
	app.Post("/logout", jwtMiddleware.Validate, logoutController.Logout)
	app.Post("/bookmark", jwtMiddleware.Validate, favouriteController.Bookmark)
	app.Post("/viewed", jwtMiddleware.Validate, viewedController.Viewed)
	app.Get("/movie/:id<int>", jwtMiddleware.Validate, movieController.Get)
	app.Get("/recommend-like/:id<int>", jwtMiddleware.Validate, itemsToItemController.Like)
	app.Get("/segments", segmentsController.GetSegmentsForUser)

	log.Fatal(app.Listen(":3002"))
}
