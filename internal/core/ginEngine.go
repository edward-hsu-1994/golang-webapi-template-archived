package core

import "github.com/gin-gonic/gin"

type GinEngine struct {
	*gin.Engine
	Container *Container
}

func NewGinEngine(engine *gin.Engine, container *Container) *GinEngine {
	instance := &GinEngine{
		Engine:    engine,
		Container: container,
	}

	return instance
}

func (this *GinEngine) UseControllers() {
	if len(this.Container.controllers) == 0 {
		return
	}

	for _, element := range this.Container.controllers {
		this.Container.Invoke(element)
	}
}
