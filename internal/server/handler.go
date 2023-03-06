package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/service"
	"log"
)

func (s *server) indexHandler(c *fiber.Ctx) error {
	t := c.Locals("user").(*jwt.Token)
	log.Println("token", t.Claims)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (s *server) authHandler(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "bad login request",
		})
	}

	errStatusCode := map[error]int{
		service.ErrUsernameMandatory: fiber.StatusBadRequest,
		service.ErrPasswordMandatory: fiber.StatusBadRequest,
		service.ErrPasswordIncorrect: fiber.StatusForbidden,
		service.ErrUserNotFound:      fiber.StatusNotFound,
		service.ErrSignedToken:       fiber.StatusInternalServerError,
	}

	user, token, err := s.userService.Authenticate(request.Username, request.Password)
	if err != nil {
		return c.Status(errStatusCode[err]).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully logged in",
		"data": fiber.Map{
			"user_id":  user.ID,
			"name":     user.Name,
			"username": user.Username,
			"role":     user.Role,
			"token":    token,
		},
	})
}
