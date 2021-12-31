package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-arch/controllers/responses"
	"golang-clean-arch/usecases/in"
)

type HashtagController struct {
	retrieve in.RetrieveHashtagInput
}

func NewHashtagController(
	retrieve *in.RetrieveHashtagInput,
) HashtagController {
	return HashtagController{
		*retrieve,
	}
}

func (controller *HashtagController) Route(app *fiber.App) {
	app.Get("/hashtag", controller.List)
}

func (controller *HashtagController) List(c *fiber.Ctx) error {
	result := controller.retrieve.GetAll()
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
