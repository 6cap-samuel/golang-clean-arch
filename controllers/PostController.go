package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang-clean-arch/entities"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/in"
)

type PostController struct {
	retrieve in.RetrievePostsInput
	create   in.CreatePostInput
}

func NewPostController(
	retrieve *in.RetrievePostsInput,
	create *in.CreatePostInput,
) PostController {
	return PostController{
		*retrieve,
		*create,
	}
}

func (controller *PostController) Route(app *fiber.App) {
	app.Post("/post", controller.Create)
	app.Get("/post", controller.List)
}

func (controller *PostController) Create(c *fiber.Ctx) error {
	var request entities.Post
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exceptions.PanicIfNeeded(err)

	controller.create.Create(request)
	return c.JSON(entities.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   request.Id,
	})
}

func (controller *PostController) List(c *fiber.Ctx) error {
	responses := controller.retrieve.GetAll()
	return c.JSON(entities.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
