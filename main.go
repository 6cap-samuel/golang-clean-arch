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
	hashtagRepository := mongo.NewHashtagRepository(database)

	retrievePost := usecases.NewRetrievePostInteractor(&postRepository)
	createPost := usecases.NewCreatePostInput(
		&postRepository,
		&hashtagRepository,
	)
	retrieveHashtag := usecases.NewRetrieveHashtagInteractor(
		&hashtagRepository,
	)

	postController := controllers.NewPostController(
		&retrievePost,
		&createPost,
	)

	hashtagController := controllers.NewHashtagController(
		&retrieveHashtag,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080, https://myfoodblog-c6e9c.web.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup Routing
	postController.Route(app)
	hashtagController.Route(app)

	port := os.Getenv("PORT")
	//port := "8081"
	err := app.Listen("0.0.0.0:" + port)

	exceptions.PanicIfNeeded(err)
}
