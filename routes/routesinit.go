package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {

	// Welcome endpoint
	app.Get("/api", Welcome)

	// User endpoints
	app.Post("/api/users", CreateUser)
	app.Get("/api/users", GetUsers)
	app.Get("/api/users/:id", GetUser)
	app.Delete("/api/users/:id", DeleteUser)

	// Product endpoints
	app.Post("/api/products", CreateProduct)
	app.Get("/api/products", GetProducts)
	app.Get("/api/products/:id", GetProduct)
	app.Put("/api/products/:id", UpdateProduct)

	// Order endpoints
	app.Post("/api/orders", CreateOrder)
	app.Get("/api/orders", GetOrders)
	app.Get("/api/orders/:id", GetOrder)

	// User Location
	app.Post("/api/userlocation", CreateUserLocation)
	app.Get("/api/userlocation", GetUserLocations)
	app.Get("/api/userlocation/:id", GetUserLocation)
	app.Put("/api/userlocation/:id", UpdateUserLocation)
	app.Delete("/api/userlocation/:id", DeleteUserLocation)
}
