package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"golang-clean-arch/entities"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(bsoncore.ValidationError)
	if ok {
		return ctx.JSON(entities.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(entities.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
