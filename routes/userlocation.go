package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonyunjae/fiber-api/database"
	"github.com/jeonyunjae/fiber-api/models"
)

func CreateResponseUserLocation(userLocation models.UserLocation) models.UserLocation {
	return models.UserLocation{ID: userLocation.ID, Lon: userLocation.Lon, Lat: userLocation.Lat, CityCode: userLocation.CityCode}
}

func CreateUserLocation(c *fiber.Ctx) error {
	var UserLocation models.UserLocation

	if err := c.BodyParser(&UserLocation); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//1. decimaltree
	//2. dictionary
	//3. kdtree
	//4. list
	//models.UserLocation = list.UserlocationList.CreateUserlocationList(UserLocation)
	//5. orm
	//6. query

	database.Database.Db.Create(&UserLocation)

	responseUserLocation := CreateResponseUserLocation(UserLocation)
	return c.Status(200).JSON(responseUserLocation)
}

func GetUserLocations(c *fiber.Ctx) error {
	userlocations := []models.UserLocation{}
	database.Database.Db.Find(&userlocations)
	responseUserLocations := []models.UserLocation{}
	for _, userlocation := range userlocations {
		responseUserLocation := CreateResponseUserLocation(userlocation)
		responseUserLocations = append(responseUserLocations, responseUserLocation)
	}

	return c.Status(200).JSON(responseUserLocations)
}

func findUserLocation(id int, userLocation *models.UserLocation) error {
	database.Database.Db.Find(&userLocation, "id = ?", id)
	if userLocation.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUserLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userLocation models.UserLocation

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUserLocation(id, &userLocation); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUserLocation := CreateResponseUserLocation(userLocation)

	return c.Status(200).JSON(responseUserLocation)
}

func UpdateUserLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userLocation models.UserLocation

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findUserLocation(id, &userLocation)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUserLocation struct {
		Lon      float64 `json:"Lon"`
		Lat      float64 `json:"Lat"`
		CityCode uint64  `json:"CityCode"`
	}

	var updateData UpdateUserLocation

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	userLocation.Lon = updateData.Lon
	userLocation.Lat = updateData.Lat
	userLocation.CityCode = updateData.CityCode

	database.Database.Db.Save(&userLocation)

	responseUserLocation := CreateResponseUserLocation(userLocation)

	return c.Status(200).JSON(responseUserLocation)

}

func DeleteUserLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userLocation models.UserLocation

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findUserLocation(id, &userLocation)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&userLocation).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}
