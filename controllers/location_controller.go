package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"project/dto/requestdto"
	"project/dto/responsedto"
	"project/models"
	"project/services"
	"project/util"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateLocation(c *fiber.Ctx) error {
	start := time.Now()
	var user models.Locations
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := services.CreateLocation(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	duration := time.Since(start).Nanoseconds()

	return c.JSON(fiber.Map{
		"status":  "success",
		"id":      user.ID,
		"time_ns": duration,
	})
}

func GetLocationByCategory(c *fiber.Ctx) error {
	start := time.Now()
	category := c.Params("category")
	location, err := services.GetLocationByCategory(category)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	duration := time.Since(start).Nanoseconds()

	return c.JSON(fiber.Map{"locations": location, "time_ns": duration})
}

func GetLocationsByCategoryAndRadius(c *fiber.Ctx) error {
	start := time.Now()
	var searchReq requestdto.LocationSearchRequest

	if err := c.BodyParser(&searchReq); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if searchReq.Category == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing field: category"})
	}

	if searchReq.Latitude == 0.0 {
		return c.Status(400).JSON(fiber.Map{"error": "Missing or invalid field: latitude"})
	}

	if searchReq.Lonitude == 0.0 {
		return c.Status(400).JSON(fiber.Map{"error": "Missing or invalid field: longitude"})
	}

	if searchReq.RadiusKm == 0.0 {
		return c.Status(400).JSON(fiber.Map{"error": "Missing or invalid field: radius_km"})
	}

	category := searchReq.Category
	lat := searchReq.Latitude
	lon := searchReq.Lonitude
	radius := searchReq.RadiusKm

	location, err := services.GetLocationsByCategoryAndRadius(category, lat, lon, radius)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	var filteredLocations []responsedto.LocationWithDistance

	for _, loc := range location {

		distance := util.Haversine(lat, lon, loc.Latitude, loc.Longitude)

		if distance <= radius {
			filteredLocations = append(filteredLocations, responsedto.LocationWithDistance{
				ID:         loc.ID,
				Name:       loc.Name,
				Address:    loc.Address,
				Latitude:   loc.Latitude,
				Longitude:  loc.Longitude,
				DistanceKm: distance,
				Category:   loc.Category,
			})
		}
	}

	duration := time.Since(start).Nanoseconds()

	return c.JSON(fiber.Map{
		"locations": filteredLocations,
		"time_ns":   duration,
	})
}

func GetTripCost(c *fiber.Ctx) error {
	start := time.Now()
	var searchReq requestdto.TripCostRequest
	locationid := c.Params("location_id")

	if err := c.BodyParser(&searchReq); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if searchReq.Latitude == 0.0 {
		return c.Status(400).JSON(fiber.Map{"error": "Missing or invalid field: latitude"})
	}

	if searchReq.Lonitude == 0.0 {
		return c.Status(400).JSON(fiber.Map{"error": "Missing or invalid field: longitude"})
	}

	lat := searchReq.Latitude
	lon := searchReq.Lonitude

	num, err := strconv.Atoi(locationid)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Location Id"})
	}

	location, err := services.GetLocationById(num)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	dto := requestdto.TollGuruRequest{
		From: requestdto.TollGuruInternal{
			Lat: lat,
			Lng: lon,
		},
		To: requestdto.TollGuruInternal{
			Lat: location.Latitude,
			Lng: location.Longitude,
		},
	}

	dtoJSON, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	totalFuel, totalToll, err := TollGuru(string(dtoJSON))

	if err != nil {
		return err
	}

	duration := time.Since(start).Nanoseconds()

	return c.JSON(fiber.Map{
		"total_cost": totalFuel + totalToll,
		"fuel_cost":  totalFuel,
		"toll_cost":  totalToll,
		"time_ns":    duration,
	})
}

type Response struct {
	Routes []struct {
		Costs struct {
			Fuel float64 `json:"fuel"`
			Tag  float64 `json:"tag"`
		} `json:"costs"`
	} `json:"routes"`
}

func TollGuru(requestbody string) (float64, float64, error) {
	url := os.Getenv("TOLLGURU_API")
	key := os.Getenv("TOLLGURU_API_KEY")

	payload := strings.NewReader(requestbody)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 0.0, 0.0, err
	}

	totalFuelCost := 0.0
	totalTollCost := 0.0

	for _, route := range response.Routes {
		totalFuelCost += route.Costs.Fuel
		totalTollCost += route.Costs.Tag
	}

	return totalFuelCost, totalTollCost, nil
}
