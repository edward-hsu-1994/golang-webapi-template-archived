package services

import (
	"golang-webapi-template/domain/errors"
	"golang-webapi-template/domain/models"
	"golang-webapi-template/domain/repositories"
)

type WeatherService struct {
	weatherRepository repositories.IWeatherRepository
}

func NewWeatherService(
	weatherRepository repositories.IWeatherRepository,
) *WeatherService {
	instance := &WeatherService{
		weatherRepository: weatherRepository,
	}

	return instance
}

func (this *WeatherService) ListWeather() []*models.Weather {
	return this.weatherRepository.FindAllWeather()
}

func (this *WeatherService) FindWeatherByLocation(location string) (*models.Weather, error) {
	temp := this.weatherRepository.FindAllWeather()

	for _, weather := range temp {
		if weather.Location == location {
			return weather, nil
		}
	}
	return nil, &errors.NotFoundLocationError{}
}
