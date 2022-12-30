package core

import (
	"github.com/gofiber/fiber/v2"
)

type FiberEngine struct {
	*fiber.App
	Container *Container
}

func NewFiberEngine(engine *fiber.App, container *Container) *FiberEngine {
	instance := &FiberEngine{
		App:       engine,
		Container: container,
	}

	return instance
}

func (this *FiberEngine) UseControllers() {
	if len(this.Container.controllers) == 0 {
		return
	}

	for _, element := range this.Container.controllers {
		this.Container.Invoke(element)
	}
}
