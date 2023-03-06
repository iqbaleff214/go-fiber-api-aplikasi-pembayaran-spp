package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/repository"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/service"
)

type Server interface {
	Run(port string) error
}

type server struct {
	app         *fiber.App
	userService service.UserService
}

func NewServer(db *sql.DB) Server {
	var out server

	out.app = fiber.New()
	out.app.Use(cors.New())
	out.userService = service.NewUserService(repository.NewUserRepository(db))

	out.SetupRoute(out.app)

	return &out
}

func (s *server) Run(port string) error {
	return s.app.Listen(port)
}
