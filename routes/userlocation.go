package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeonyunjae/fiber-api/models"
	"github.com/jeonyunjae/fiber-api/service"
)

func CreateResponsePositionAddressInfo(positionAddressInfo models.PositionAddressInfo) models.PositionAddressInfo {
	return models.PositionAddressInfo{ID: positionAddressInfo.ID, Lon: positionAddressInfo.Lon, Lat: positionAddressInfo.Lat}
}

func PositionAddressInfoInsert(c *fiber.Ctx) error {
	var PositionAddressInfo models.PositionAddressInfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.ServiceInsert(PositionAddressInfo)

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

func PositionAddressInfoReads(c *fiber.Ctx) error {
	var PositionAddressInfo models.PositionAddressInfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.ServiceRead(PositionAddressInfo)

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

func PositionAddressInfoUpdate(c *fiber.Ctx) error {
	var PositionAddressInfo models.PositionAddressInfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.ServiceUpdate(PositionAddressInfo)

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

func PositionAddressInfoDelete(c *fiber.Ctx) error {
	var PositionAddressInfo models.PositionAddressInfo

	if err := c.BodyParser(&PositionAddressInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.ServiceDelete(PositionAddressInfo)

	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
	return c.Status(200).JSON(responsePositionAddressInfo)
}

// func GetPositionAddressInfos(c *fiber.Ctx) error {
// 	PositionAddressInfos := []models.PositionAddressInfo{}
// 	gorm.Database.Db.Find(&PositionAddressInfos)
// 	responsePositionAddressInfos := []models.PositionAddressInfo{}
// 	for _, PositionAddressInfo := range PositionAddressInfos {
// 		responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)
// 		responsePositionAddressInfos = append(responsePositionAddressInfos, responsePositionAddressInfo)
// 	}

// 	return c.Status(200).JSON(responsePositionAddressInfos)
// }

// func findPositionAddressInfo(id int, PositionAddressInfo *models.PositionAddressInfo) error {
// 	gorm.Database.Db.Find(&PositionAddressInfo, "id = ?", id)
// 	if PositionAddressInfo.ID == 0 {
// 		return errors.New("user does not exist")
// 	}
// 	return nil
// }

// func GetPositionAddressInfo(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var PositionAddressInfo models.PositionAddressInfo

// 	if err != nil {
// 		return c.Status(400).JSON("Please ensure that :id is an integer")
// 	}

// 	if err := findPositionAddressInfo(id, &PositionAddressInfo); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)

// 	return c.Status(200).JSON(responsePositionAddressInfo)
// }

// func UpdatePositionAddressInfo(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var PositionAddressInfo models.PositionAddressInfo

// 	if err != nil {
// 		return c.Status(400).JSON("Please ensure that :id is an integer")
// 	}

// 	err = findPositionAddressInfo(id, &PositionAddressInfo)

// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	type UpdatePositionAddressInfo struct {
// 		Lon      float64 `json:"Lon"`
// 		Lat      float64 `json:"Lat"`
// 		CityCode uint64  `json:"CityCode"`
// 	}

// 	var updateData UpdatePositionAddressInfo

// 	if err := c.BodyParser(&updateData); err != nil {
// 		return c.Status(500).JSON(err.Error())
// 	}

// 	PositionAddressInfo.Lon = updateData.Lon
// 	PositionAddressInfo.Lat = updateData.Lat
// 	PositionAddressInfo.CityCode = updateData.CityCode

// 	gorm.Database.Db.Save(&PositionAddressInfo)

// 	responsePositionAddressInfo := CreateResponsePositionAddressInfo(PositionAddressInfo)

// 	return c.Status(200).JSON(responsePositionAddressInfo)

// }

// func DeletePositionAddressInfo(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var PositionAddressInfo models.PositionAddressInfo

// 	if err != nil {
// 		return c.Status(400).JSON("Please ensure that :id is an integer")
// 	}

// 	err = findPositionAddressInfo(id, &PositionAddressInfo)

// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	if err = gorm.Database.Db.Delete(&PositionAddressInfo).Error; err != nil {
// 		return c.Status(404).JSON(err.Error())
// 	}
// 	return c.Status(200).JSON("Successfully deleted User")
// }
