package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"

)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func setupRoutes(app *fiber.App) {
	app.Get("/api/movies/", getMovies)
	app.Get("/api/movies/:id/", getMovie)
	app.Post("/api/movies/:id/", createMovie)
	app.Put("/api/movies/:id/", updateMovie)
	app.Delete("/api/movies/:id/", deleteMovie)
}

func main() {
	movies = append(movies, Movie{
		ID: "1", 
		Isbn: "625738",
		Title: "Movie one",
		Director: &Director{
			Firstname: "John",
			Lastname: "Doe",
		}})

	movies = append(movies, Movie{
		ID: "2", 
		Isbn: "936738",
		Title: "Movie two",
		Director: &Director{
			Firstname: "Steve",
			Lastname: "Smith",
		}})

	app := fiber.New()
	setupRoutes(app)
	fmt.Printf("Starting server at 8000...")
	log.Fatal(app.Listen(":8000"))

}