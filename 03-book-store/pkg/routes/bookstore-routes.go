package routes

import (
	"github.com/gofiber/fiber/v2"
	"book_store/pkg/controllers"
)

func RegisterBookStoreRoutes(app *fiber.App) {
	app.Post("/book", controllers.CreateBook)
	app.Get("/book", controllers.GetBook)
	app.Get("/book/:bookId", controllers.GetBookById)
	app.Put("/book/:bookId", controllers.UpdateBook)
	app.Delete("/book/:bookId", controllers.DeleteBook)
}