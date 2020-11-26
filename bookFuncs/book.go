package bookFuncs

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Book struct {
	BookId     int
	BookName   string
	BookAuthor string
}

var books = []Book{
	Book{
		1,
		"Kaustubh",
		"Kishore",
	},
	Book{
		2,
		"Revolta",
		"Go",
	},
}

func GetAllBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func GetOneBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	found := false

	for i := 0; i < len(books); i++ {
		if books[i].BookId == id {
			found = true
			return c.Status(200).JSON(books[i])
		}
	}

	if !found {
		c.SendStatus(404)
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "Not Found",
		})
	}
	return nil
}

func AddBook(c *fiber.Ctx) error {
	bookTemp := new(Book)
	if err := c.BodyParser(bookTemp); err != nil {
		return c.Status(400).JSON(err)
	}
	books = append(books, *bookTemp)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	found := false
	id, _ := strconv.Atoi(c.Params("id"))
	for i := 0; i < len(books); i++ {
		if books[i].BookId == id {
			bookTemp := Book(books[i])
			found = true
			books = append(books[:i], books[i+1:]...)
			return c.Status(200).JSON(bookTemp)
		}
	}

	if !found {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"Error":   "Book not found",
		})
	}
	return nil
}
