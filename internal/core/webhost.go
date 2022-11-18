package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type WebHost struct {
	container          *Container
	startupConstructor any
}

func NewWebHost(startupConstructor any) *WebHost {
	instance := &WebHost{
		container:          NewContainer(dig.New()),
		startupConstructor: startupConstructor,
	}

	return instance
}

func (this *WebHost) Run() error {
	this.container.Provide(getConfiguration)

	this.container.Provide(this.startupConstructor)

	err := this.container.Invoke(func(startup *Startup) error {
		err := startup.Bootstrap(this)
		return err
	})

	if err != nil {
		return err
	}

	err = this.container.Invoke(func(engine *gin.Engine) error {
		return engine.Run()
	})

	if err != nil {
		return err
	}

	return nil
}

// ref. https://tutorialedge.net/golang/parsing-json-with-golang/
func getConfiguration() Configuration {
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
}
