package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/repositories/mongo"
	"golang-clean-arch/usecases"
	"os"
)

func main() {
	configuration := configurations.New()
	database := configurations.NewMongoDatabase(configuration)

	postRepository := mongo.NewPostRepository(database)
	storeRepository := mongo.NewStoreRepository(database)

	retrievePost := usecases.NewRetrievePostInteractor(&postRepository, &storeRepository)
	createPost := usecases.NewCreatePostInput(&postRepository, &storeRepository)

	postController := controllers.NewPostController(
		&retrievePost,
		&createPost,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup Routing
	postController.Route(app)

	port := os.Getenv("PORT")
	err := app.Listen("0.0.0.0:" + port)

	exceptions.PanicIfNeeded(err)
}
