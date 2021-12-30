package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-arch/controllers/mappers"
	"golang-clean-arch/controllers/requests"
	"golang-clean-arch/controllers/responses"
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
	var request requests.CreatePostRequest
	err := c.BodyParser(&request)
	var storeEntity = mappers.CreatePostRequestToPostMapper(request)

	exceptions.PanicIfNeeded(err)

	controller.create.Create(*storeEntity)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   storeEntity.Id,
	})
}

func (controller *PostController) List(c *fiber.Ctx) error {
	result := controller.retrieve.GetAll()
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
