package rest

import (
	"antrein/dd-dashboard-auth/model/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg *config.Config, app *fiber.App) error {
	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Rest.Port))
}
