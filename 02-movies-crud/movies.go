package main

import (
	"math/rand"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func getMovies(c *fiber.Ctx) error {
	return c.JSON(movies)
}

func getMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, item := range movies {
		if item.ID == id {
			return c.JSON(item)
		}
	}
	return c.Status(404).JSON(fiber.Map{"message": "Movie not found"})
}

func createMovie(c *fiber.Ctx) error {
	movie := new(Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, *movie)
	return c.JSON(movie)
}

func updateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie
			_ = c.BodyParser(&movie)
			movie.ID = id
			movies = append(movies, movie)
			return c.JSON(movie)
		}
	}
	return c.Status(404).JSON(fiber.Map{"message":"Movie not found"})
}

func deleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index + 1:]...)
			return c.JSON(fiber.Map{"message": "Movie deleted successfully"})
		}
	}
	return c.Status(404).JSON(fiber.Map{"message":"Movie not found"})
}