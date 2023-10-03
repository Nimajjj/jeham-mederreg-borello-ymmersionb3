package main

import (
	"fmt"

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

    pebbles := LoadJSONPebbles()

    pebble_repo := repo.NewPebbleRepo()
    fmt.Println("----pebbles----")
    for _, pebble := range pebbles {
        pebble_repo.InsertNewPebble(&pebble)
    }
    fmt.Println("---------------")

	// RUN
	routes.Run(HOST) // run api
}
