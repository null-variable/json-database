package main

import (
	"github.com/gofiber/fiber/v2"
	utils "github.com/th-release/utils/go"
)

func SetupRoutes(app *fiber.App) *fiber.App {
	app.Get("/", ServerStatusHandler)
	app.Get("/status", ServerStatusHandler)
	app.Post("/create/table", CreateTableHandler)
	return app
}

func ServerStatusHandler(c *fiber.Ctx) error {
	response := utils.BasicResponse{
		Success: true,
		Message: nil,
		Data:    nil,
	}

	return c.JSON(response)
}

func CreateTableHandler(c *fiber.Ctx) error {
	return c.SendString("asd")
}
