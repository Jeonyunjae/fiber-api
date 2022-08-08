package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonyunjae/fiber-api/database"
	"github.com/jeonyunjae/fiber-api/routes"
	"github.com/jeonyunjae/fiber-api/service"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	routes.RoutesInit(app)
	service.ServiceInit()

	log.Fatal(app.Listen(":3000"))

}
