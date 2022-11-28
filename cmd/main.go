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

	webhost := core.NewWebHost(application.NewStartup)
	err := webhost.Run()

	if err != nil {
		fmt.Printf("Application error: %s", err.Error())
	}
}
