package main

import (
	"fmt"
	"golang-webapi-template/application"
	"golang-webapi-template/core"
)

// @title           XXX API
// @version         1.0
// @description     API template
func main() {
	fmt.Println("Starting...")

	host, err := core.NewWebHostBuilder().UseStartup(application.NewStartup).UseGinEngine().Build()

	if err != nil {
		fmt.Printf("Application initial: %s", err.Error())
	}

	err = host.Run()

	if err != nil {
		fmt.Printf("Application runtime error: %s", err.Error())
	}
}
