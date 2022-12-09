package models

import "golang-webapi-template/domain/models"

type WeatherResponse struct {
	Location     string `json:"location"`
	TemperatureC string `json:"temperatureC`
}

func WeatherToResponse(weather *models.Weather) *WeatherResponse {
	return &WeatherResponse{
		Location:     weather.Location,
		TemperatureC: weather.TemperatureC,
	}
}
