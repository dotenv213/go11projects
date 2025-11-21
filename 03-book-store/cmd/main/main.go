package main

import (
	"book_store/pkg/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.RegisterBookStoreRoutes(app)
	fmt.Printf("Running server...")
	log.Fatal(app.Listen(":9010"))
}
