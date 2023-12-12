package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Run() error {
	app := fiber.New()
	setupRoutes(app)
	httpPort := os.Getenv("PORT")
	if err := app.Listen(":" + httpPort); err != nil {
		return err
	}
	return nil
}

func setupRoutes(app *fiber.App) {
	app.Get("/", getBlockchain)
	app.Post("/", writeBlock)
}
