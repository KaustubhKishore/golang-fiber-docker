package main

import (
	"fmt"
	"github.com/KaustubhKishore/fiberdemo/bookFuncs"
	"github.com/gofiber/fiber/v2"
)

func setupAllRoutes(app *fiber.App) {
	app.Get("/", bookFuncs.GetAllBooks)
	app.Get("/all", bookFuncs.GetAllBooks)
	app.Get("/api/getonebook/:id", bookFuncs.GetOneBook)
	app.Post("/api/addonebook", bookFuncs.AddBook)
	app.Delete("/api/deletebook/:id", bookFuncs.DeleteBook)
}

func main() {
	app := fiber.New()
	setupAllRoutes(app)

	fmt.Println("Running...")
	app.Listen("0.0.0.0:3000")
}
