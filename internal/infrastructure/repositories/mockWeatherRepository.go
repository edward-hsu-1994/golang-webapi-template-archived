package repositories

import (
	"golang-webapi-template/domain/models"
	"golang-webapi-template/domain/repositories"
)

type MockWeatherRepository struct {
}

func NewMockWeatherRepository() repositories.IWeatherRepository {
	instance := &MockWeatherRepository{}

	return instance
}

func (this *MockWeatherRepository) FindAllWeather() []*models.Weather {
	data := []*models.Weather{
		&models.Weather{
			Location:     "Taipei",
			TemperatureC: "28",
			Date:         "2022-12-10",
		},
		&models.Weather{
			Location:     "Tainan",
			TemperatureC: "29",
			Date:         "2022-12-10",
		},
	}

	return data
}
