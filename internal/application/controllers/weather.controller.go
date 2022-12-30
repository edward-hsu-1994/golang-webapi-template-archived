package controllers

import (
	r "golang-webapi-template/application/models"
	"golang-webapi-template/core"
	"golang-webapi-template/domain/models"
	"golang-webapi-template/domain/services"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/gofiber/fiber/v2"
)

type WeatherController struct {
	weatherService *services.WeatherService
}

func NewWeatherController(
	router *core.FiberEngine,
	weatherService *services.WeatherService,
) *WeatherController {
	instance := &WeatherController{
		weatherService: weatherService,
	}

	router.Get("/api/weather", instance.GetWeathers)
	router.Get("/api/weather/:uid\\:myAction", instance.GetWeathers)

	return instance
}

// @Summary	Find all weather info
// @Tags	Weathers
// @version	1.0
// @produce	application/json
// @Success	200	{array}	models.Weather
// @Router	/api/weather	[get]
func (this *WeatherController) GetWeathers(ctx *fiber.Ctx) error {
	var data []*r.WeatherResponse

	From(this.weatherService.ListWeather()).SelectT(func(item *models.Weather) *r.WeatherResponse {
		return r.WeatherToResponse(item)
	}).ToSlice(&data)

	ctx.JSON(data)

	return ctx.SendStatus(200)
}
