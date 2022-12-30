package application

import (
	"golang-webapi-template/application/controllers"
	"golang-webapi-template/application/middlewares"
	"golang-webapi-template/core"
	"golang-webapi-template/domain/services"
	"golang-webapi-template/infrastructure/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
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

func Configure(engine *core.FiberEngine) {
	// Use json logging
	engine.Use(middlewares.Log)

	// Add static file to routes
	engine.Static("/", "./assets", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	engine.UseControllers()
}
