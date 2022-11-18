package controllers

import (
	"golang-webapi-template/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherController struct {
}

func NewWeatherController(
	router *gin.Engine,
) *WeatherController {
	instance := &WeatherController{}

	router.GET("/api/weather", instance.GetWeathers)
	return instance
}

// @Summary	Find all weather info
// @Tags	Weathers
// @version	1.0
// @produce	application/json
// @Success	200	{object}	appModels.User
// @Router	/api/weather	[get]
func (this *WeatherController) GetWeathers(context *gin.Context) {
	data := []*models.Weather{
		&models.Weather{
			Location:     "Taipei",
			TemperatureC: "28",
		},
		&models.Weather{
			Location:     "Tainan",
			TemperatureC: "29",
		},
	}

	context.JSON(http.StatusOK, data)
}
