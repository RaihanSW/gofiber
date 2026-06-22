package main

import (
	"gofiber/database"
	"gofiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to API")
}

func setupRoutes(app *fiber.App) {
	//Welcome Endpoint
	app.Get("/api", welcome)

	//User Endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Get("/api/users", routes.GetUsersList)
	app.Get("/api/users/:id", routes.GetUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//Product Endpoint
	app.Post("/api/products", routes.CreateProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Get("/api/products", routes.GetProductsList)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	// Order Endpoint
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrdersList)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {
	database.ConnectToDB()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
