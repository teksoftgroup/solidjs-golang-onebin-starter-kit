package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/teksoftgroup/embed-solidjs/client"
)

func main() {
	app := fiber.New()

	app.Get("/hello.json", handleGetJSON)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         client.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func handleGetJSON(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello from server",
	})
}
