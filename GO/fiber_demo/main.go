package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/programmerug/fibergorm/config"
	"github.com/programmerug/fibergorm/handlers"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/:id", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/:id", handlers.UpdateDog)
	app.Delete("/dogs/:id", handlers.RemoveDog)

	log.Fatal(app.Listen(":3000"))
}
