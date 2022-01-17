package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-arch/controllers/responses"
	"golang-clean-arch/usecases/in"
)

type AuthController struct {
	get in.GetGoogleSSOLinkWithEmailinput
}

func NewAuthController(
	getEmailLink *in.GetGoogleSSOLinkWithEmailinput,
) AuthController {
	return AuthController{
		*getEmailLink,
	}
}

func (controller *AuthController) Route(app *fiber.App) {
	app.Get("/auth/email", controller.GetEmailLink)
}

func (controller *AuthController) GetEmailLink(c *fiber.Ctx) error {
	result := controller.get.Get(c.Get("email"))
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
