package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service"
)

func CreateResponsePositionAddressInfo(positionAddressInfo models.Positionaddressinfo) models.Positionaddressinfo {
	return models.Positionaddressinfo{Usercode: positionAddressInfo.Usercode, Loclongtitude: positionAddressInfo.Loclongtitude, Loclatitude: positionAddressInfo.Loclatitude}
}

func CreateResponsePositionAddressDistanceInfo(positionAddressDistanceInfo models.PositionaddressDistanceInfo) models.PositionaddressDistanceInfo {
	return models.PositionaddressDistanceInfo{Usercode: positionAddressDistanceInfo.Usercode, Loclongtitude: positionAddressDistanceInfo.Loclongtitude, Loclatitude: positionAddressDistanceInfo.Loclatitude, Distance: positionAddressDistanceInfo.Distance, Count: positionAddressDistanceInfo.Count}
}

func PositionAddressInfoInsert(c *fiber.Ctx) error {
	var PositionAddressInfo models.Positionaddressinfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := service.ServiceInsert(PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

func PositionAddressInfoReadsLimit(c *fiber.Ctx) error {
	var PositionaddressDistanceInfo models.PositionaddressDistanceInfo
	var PositionaddressDistanceInfos []models.PositionaddressDistanceInfo
	var err error

	if err = c.BodyParser(&PositionaddressDistanceInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if PositionaddressDistanceInfos, err = service.ServiceReadsLimit(PositionaddressDistanceInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(PositionaddressDistanceInfos)
}

func PositionAddressInfoRead(c *fiber.Ctx) error {
	var PositionAddressInfo models.Positionaddressinfo
	var result map[string]models.Positionaddressinfo
	var err error

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if result, err = service.ServiceRead(PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}

func PositionAddressInfoUpdate(c *fiber.Ctx) error {
	var PositionAddressInfo models.Positionaddressinfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := service.ServiceUpdate(PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

func PositionAddressInfoDelete(c *fiber.Ctx) error {
	var PositionAddressInfo models.Positionaddressinfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := service.ServiceDelete(PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}
