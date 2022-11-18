package main

import (
	"fmt"
	"golang-webapi-template/app"
	"golang-webapi-template/core"
)

func main() {
	fmt.Println("Starting...")

	webhost := core.NewWebHost(app.NewStartup)
	err := webhost.Run()

	if err != nil {
		fmt.Printf("Application error: %s", err.Error())
	}
}
