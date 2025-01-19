package services

import (
	"errors"
	"project/database"
	"project/models"
)

func CreateLocation(user *models.Locations) error {
	return database.DB.Create(user).Error
}

func GetLocationByCategory(category string) ([]models.Locations, error) {
	var locations []models.Locations
	result := database.DB.Where("category = ?", category).Find(&locations)
	if result.RowsAffected == 0 {
		return locations, errors.New("no data found")
	}
	return locations, result.Error
}

func GetLocationById(id int) (models.Locations, error) {
	var locations models.Locations
	result := database.DB.Where("id = ?", id).First(&locations)
	if result.RowsAffected == 0 {
		return locations, errors.New("no location with given id")
	}
	return locations, result.Error
}

func GetLocationsByCategoryAndRadius(category string, lat float64, lon float64, radius float64) ([]models.Locations, error) {
	var locations []models.Locations
	result := database.DB.Where("category = ?", category).Find(&locations)

	return locations, result.Error
}
