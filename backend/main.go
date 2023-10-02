package main

import (
	"prgc/repo"
	"prgc/routes"
)

const HOST = "localhost:8080"

func main() {
	// INIT
	repo.Init()   // repo (db) setup
	routes.Init() // gin setup

	// MAIN
	routes.SetupRoutes() // init each routes

	// RUN
	routes.Run(HOST) // run api
}
