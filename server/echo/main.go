package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teksoftgroup/embed-solidjs/client"
)

func main() {
	app := echo.New()

	app.GET("/hello.json", handleGetJSON)

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: client.BuildHTTPFS(),
		HTML5:      true,
	}))
	log.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func handleGetJSON(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello from server",
	})
}
