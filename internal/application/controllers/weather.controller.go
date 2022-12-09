package controllers

import (
	r "golang-webapi-template/application/models"
	"golang-webapi-template/domain/models"
	"golang-webapi-template/domain/services"
	"net/http"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
)

type WeatherController struct {
	weatherService *services.WeatherService
}

func NewWeatherController(
	router *gin.Engine,
	weatherService *services.WeatherService,
) *WeatherController {
	instance := &WeatherController{
		weatherService: weatherService,
	}

	router.GET("/api/weather", instance.GetWeathers)
	return instance
}

// @Summary	Find all weather info
// @Tags	Weathers
// @version	1.0
// @produce	application/json
// @Success	200	{array}	models.Weather
// @Router	/api/weather	[get]
func (this *WeatherController) GetWeathers(context *gin.Context) {
	var data []*r.WeatherResponse

	From(this.weatherService.ListWeather()).SelectT(func(item *models.Weather) *r.WeatherResponse {
		return r.WeatherToResponse(item)
	}).ToSlice(&data)

	context.JSON(http.StatusOK, data)
}
