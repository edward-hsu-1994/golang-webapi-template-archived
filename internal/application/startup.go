package application

import (
	"golang-webapi-template/application/controllers"
	"golang-webapi-template/application/middlewares"
	"golang-webapi-template/core"
	"golang-webapi-template/domain/services"
	"golang-webapi-template/infrastructure/repositories"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var configuration *core.Configuration

func NewStartup(config core.Configuration) *core.WebHostStartup {
	instance := &core.WebHostStartup{}

	instance.ConfigureServicesFunc = ConfigureServices
	instance.ConfigureFunc = Configure

	configuration = &config

	return instance
}

func ConfigureServices(container *core.Container) {
	// Add controller's constructor into container
	container.AddControllers(controllers.NewWeatherController)

	container.Provide(services.NewWeatherService)
	container.Provide(repositories.NewMockWeatherRepository)
}

func Configure(engine *core.GinEngine) {
	// Use json logging
	engine.Use(middlewares.JsonLoggerMiddleware())
	engine.Use(gin.Recovery())

	// Add static file to routes
	engine.Use(static.Serve("/", static.LocalFile("./assets", false)))

	engine.UseControllers()
}
