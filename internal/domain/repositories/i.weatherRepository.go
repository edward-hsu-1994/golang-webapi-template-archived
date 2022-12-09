package repositories

import "golang-webapi-template/domain/models"

type IWeatherRepository interface {
	FindAllWeather() []*models.Weather
}
