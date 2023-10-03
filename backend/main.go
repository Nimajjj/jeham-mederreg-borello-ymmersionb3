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
    fmt.Println("---pebbles---")
    for _, p := range pebbles{
        fmt.Printf("%d -> %s\n", p.ID, p.Title)
    }

	// RUN
	routes.Run(HOST) // run api
}
