package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"book_store/pkg/models"
)

func GetBook(c *fiber.Ctx) error {
	newBooks := models.GetAllBooks()
	return c.JSON(newBooks)
}

func GetBookById(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	bookDetails, _ := models.GetBookById(ID)
	return c.JSON((bookDetails))
}

func CreateBook(c *fiber.Ctx) error {
	createBook := &models.Book{}
	if err := c.BodyParser(createBook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	b := createBook.CreateBook()
	return c.JSON(b)
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	book := models.DeleteBook(ID)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	bookDetails, db := models.GetBookById(ID)
    
	updateBook := &models.Book{}
	if err := c.BodyParser(updateBook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Data"})
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	return c.JSON(bookDetails)
}

