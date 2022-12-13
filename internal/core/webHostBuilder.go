package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type IHostBuilder interface {
	Build() IHost
}

type IHost interface {
	Run() error
}

type WebHostBuilder struct {
	container    *Container
	provideFuncs []any
	runFunc      any
}

func NewWebHostBuilder() *WebHostBuilder {
	var builder = &WebHostBuilder{
		container: NewContainer(dig.New()),
	}

	// #region init default DI
	builder.container.Provide(func() *Container {
		return builder.container
	})

	builder.container.Provide(func() *dig.Container {
		return builder.container.Container
	})

	builder.container.Provide(func() Configuration {
		var configuration Configuration

		// Open our jsonFile
		jsonFile, err := os.Open("./configs/appsettings.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println("Not found configuration file: configs/appsettings.json")
			configuration = make(Configuration)
			return configuration
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal([]byte(byteValue), &configuration)

		return configuration
	})
	// #endregion

	return builder
}

func (this *WebHostBuilder) UseStartup(startupConstructor any) *WebHostBuilder {
	this.provideFuncs = append(this.provideFuncs, startupConstructor)
	return this
}

func (this *WebHostBuilder) UseConfiguration(path *string) *WebHostBuilder {
	var filePath = "./configs/appsettings.json"
	if path != nil {
		filePath = *path
	}
	this.provideFuncs = append(this.provideFuncs, func() Configuration {
		var configuration Configuration

		// Open our jsonFile
		jsonFile, err := os.Open(filePath)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println("Not found configuration file: " + filePath)
			configuration = make(Configuration)
			return configuration
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal([]byte(byteValue), &configuration)

		return configuration
	})

	return this
}

func (this *WebHostBuilder) UseGinEngine() *WebHostBuilder {
	this.provideFuncs = append(this.provideFuncs, func() *gin.Engine {
		return gin.New()
	})
	this.provideFuncs = append(this.provideFuncs, func(engine *gin.Engine, container *Container) *GinEngine {
		return NewGinEngine(engine, container)
	})

	this.runFunc = func(engine *gin.Engine) error {
		return engine.Run()
	}

	return this
}

func (this *WebHostBuilder) Build() (IHost, error) {
	var host = &WebHost{
		container: this.container,
		runFunc:   this.runFunc,
	}

	for _, provideFunc := range this.provideFuncs {
		host.container.Provide(provideFunc)
	}

	err := host.container.Invoke(func(startup *WebHostStartup, container *Container) {
		container.Invoke(startup.ConfigureServicesFunc)
	})

	if err != nil {
		return nil, err
	}

	err = host.container.Invoke(func(startup *WebHostStartup, container *Container) {
		container.Invoke(startup.ConfigureFunc)
	})

	if err != nil {
		return nil, err
	}

	return host, nil
}
