package main

import (
	"test-b/config"
	"test-b/middlewares"
	"test-b/routes"
)

// function main
func main() {
	config.InitDB()
	e := routes.New()
	// take notes http method activity
	middlewares.LogMiddlewares(e)
	// start on port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
