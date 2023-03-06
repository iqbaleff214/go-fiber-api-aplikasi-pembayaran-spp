package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwt "github.com/gofiber/jwt/v2"
	gojwt "github.com/golang-jwt/jwt/v4"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/config"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/entity"
)

func (s *server) SetupRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", s.indexHandler)
	api.Post("/login", s.authHandler)
}

// middleware for JWT auth
func (s *server) middlewareJWTHandler() func(c *fiber.Ctx) error {
	return jwt.New(jwt.Config{
		SigningKey: []byte(config.NewConfig().SecretJWT()),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  "error",
					"message": "Missing or malformed JWT",
					"data":    nil,
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid or expired JWT",
				"data":    nil,
			})
		},
	})
}

func (s *server) middlewareAdminGuardHandler(c *fiber.Ctx) error {
	t := c.Locals("user").(*gojwt.Token)
	claims := t.Claims.(gojwt.MapClaims)

	if claims["role"] == string(entity.Administrator) {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "you're not allowed to access this resource",
	})
}
