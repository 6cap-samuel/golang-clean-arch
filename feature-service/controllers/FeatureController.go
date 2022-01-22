package controllers

import (
	"feature-service/controllers/responses"
	"feature-service/usecases/in"
	"github.com/gofiber/fiber/v2"
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
