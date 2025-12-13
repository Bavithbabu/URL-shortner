package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bavith/Url_shortern/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {

	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("api/v1/urls", routes.ListUserURLs)             // New routes
	app.Get("/api/v1/url/:shortcode", routes.GetURLDetails) // get the single url detailsi
	app.Delete("/api/v1/url/:shortcode", routes.DeleteURL)  // Route for deleting the url
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
