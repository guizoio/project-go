package handler

import (
	"awesomeProject1/src/use_case/user/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(Service service.UserService) UserHandler {
	return UserHandler{Service}
}

func (u UserHandler) List(c *fiber.Ctx) error {
	result, err := u.Service.List()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ERROR")
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
