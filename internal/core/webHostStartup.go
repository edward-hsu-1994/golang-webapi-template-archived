package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// WebHostStartup requires `ConfigureServices` and `Configure`
type WebHostStartup struct {
	ConfigureServicesFunc any
	ConfigureFunc         any
}

func (this *WebHostStartup) Boot(webhost *WebHost) error {
	// #region init default DI
	webhost.container.Provide(func() *Container {
		return webhost.container
	})

	webhost.container.Provide(func() *dig.Container {
		return webhost.container.Container
	})
	// #endregion

	err := webhost.container.Invoke(this.ConfigureServicesFunc)

	if err != nil {
		return nil
	}

	webhost.container.Provide(func() *gin.Engine {
		return gin.New()
	})
	webhost.container.Provide(func(engine *gin.Engine, container *Container) *GinEngine {
		return NewGinEngine(engine, container)
	})

	err = webhost.container.Invoke(this.ConfigureFunc)

	if err != nil {
		return nil
	}

	return nil
}
