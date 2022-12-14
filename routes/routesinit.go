package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {

	// Welcome endpoint
	app.Get("/api", Welcome)

	// User Location
	app.Post("/api/PositionAddressInfoInsert", PositionAddressInfoInsert)
	app.Post("/api/PositionAddressInfoRead", PositionAddressInfoRead)
	app.Post("/api/PositionAddressInfoReadsLimit", PositionAddressInfoReadsLimit)
	app.Post("/api/PositionAddressInfoUpdate", PositionAddressInfoUpdate)
	app.Post("/api/PositionAddressInfoDelete", PositionAddressInfoDelete)
	// app.Get("/api/PositionAddressInfo", GetPositionAddressInfos)
	// app.Get("/api/PositionAddressInfo/:id", GetPositionAddressInfo)
	// app.Put("/api/PositionAddressInfo/:id", UpdatePositionAddressInfo)
	// app.Delete("/api/PositionAddressInfo/:id", DeletePositionAddressInfo)
}
