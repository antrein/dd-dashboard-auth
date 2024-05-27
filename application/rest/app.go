package rest

import (
	"antrein/dd-dashboard-auth/application/common/resource"
	"antrein/dd-dashboard-auth/application/common/usecase"
	"antrein/dd-dashboard-auth/internal/handler/rest/auth"
	"antrein/dd-dashboard-auth/model/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ApplicationDelegate(cfg *config.Config, uc *usecase.CommonUsecase, rsc *resource.CommonResource) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "BC Dashboard",
	})

	// setup gzip
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// setup cors
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	dashboard := app.Group("bc/dashboard")

	dashboard.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Makan nasi pagi-pagi, ngapain kamu disini?")
	})

	dashboard.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong!")
	})

	// routes

	// auth
	authRoute := auth.New(cfg, uc.AuthUsecase, rsc.Vld)
	authRoute.RegisterRoute(app)

	return app, nil
}
