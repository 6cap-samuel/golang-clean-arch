package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-arch/controllers/responses"
	"golang-clean-arch/usecases/in"
)

type FeatureController struct {
	retrieve in.RetrieveFeatureInput
}

func NewFeatureController(
	retrieve *in.RetrieveFeatureInput,
) FeatureController {
	return FeatureController{
		*retrieve,
	}
}

func (controller *FeatureController) Route(app *fiber.App) {
	app.Get("/features", controller.List)
}

func (controller *FeatureController) List(c *fiber.Ctx) error {
	result := controller.retrieve.GetAll()
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
