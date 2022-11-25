package main

import (
	"fmt"
	"golang-webapi-template/application"
	"golang-webapi-template/core"
)

func main() {
	fmt.Println("Starting...")

	webhost := core.NewWebHost(application.NewStartup)
	err := webhost.Run()

	if err != nil {
		fmt.Printf("Application error: %s", err.Error())
	}
}
