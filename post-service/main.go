package main

import (
	"fmt"
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/repositories/client"
	"golang-clean-arch/repositories/mongo"
	"golang-clean-arch/usecases"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("Setting up datasource")
	configuration := configurations.New()
	database := configurations.NewMongoDatabase(configuration)
	firebaseAuth := configurations.NewFirebaseAuth()

	fmt.Println("Setting up repositories")
	postRepository := mongo.NewPostRepository(database)
	hashtagRepository := mongo.NewHashtagRepository(database)
	featureRepository := mongo.NewFeatureRepository(database)
	firebaseClient := client.NewFirebaseClient(firebaseAuth)

	fmt.Println("Setting up use cases")
	retrievePost := usecases.NewRetrievePostInteractor(&postRepository)
	createPost := usecases.NewCreatePostInput(
		&postRepository,
		&hashtagRepository,
	)
	updatePost := usecases.NewUpdatePostIntractor(&postRepository)
	retrieveHashtag := usecases.NewRetrieveHashtagInteractor(
		&hashtagRepository,
	)
	detailsPost := usecases.NewRetrievePostDetailsInteractor(
		&postRepository,
	)
	getAllFeature := usecases.NewGetAllFeatureDataInteractor(
		&featureRepository,
	)
	getGoogleSsoLink := usecases.NewGetGoogleSsoLinkWithEmailInteractor(&firebaseClient)

	fmt.Println("Setting up controllers")
	postController := controllers.NewPostController(
		&retrievePost,
		&createPost,
		&updatePost,
		&detailsPost,
	)
	hashtagController := controllers.NewHashtagController(
		&retrieveHashtag,
	)
	featureController := controllers.NewFeatureController(
		&getAllFeature,
	)
	authController := controllers.NewAuthController(
		&getGoogleSsoLink,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Singapore",
	}))

	// Setup Routing
	postController.Route(app)
	hashtagController.Route(app)
	featureController.Route(app)
	authController.Route(app)

	port := os.Getenv("PORT")
	err := app.Listen("0.0.0.0:" + port)

	exceptions.PanicIfNeeded(err)
}
