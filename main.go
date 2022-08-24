package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonyunjae/fiber-api/datatype/gorm"
	"github.com/jeonyunjae/fiber-api/datatype/query"
	"github.com/jeonyunjae/fiber-api/routes"
	"github.com/jeonyunjae/fiber-api/service"
)

func main() {
	gorm.ConnectDb()
	query.ConnectDb()

	app := fiber.New()

	routes.RoutesInit(app)
	service.ServiceInit()

	log.Fatal(app.Listen(":3000"))
}
