package configurations

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-arch/exceptions"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}
