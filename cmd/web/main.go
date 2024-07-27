package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers"
	"dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

// AppConfig holds the application config

const portNumber = ":8084"

var app config.AppConfig

func main() {
	app.InProduction = false
	app.SideNav = false
	app.MainTable = false
	app.ViewCache = make(map[string]constants.ViewInteface)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	var programLevel = new(slog.LevelVar) // Info by default
	programLevel.Set(slog.LevelDebug)

	initRouteHandlers()
	initApp()

	fmt.Printf("Starting application on port %s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func initApp() {
	// init the app
	database.Init()
	// sidenav.AppSideNav.InitDropdownData()

	// queueSize := 100
	// bus := messagebus.New(queueSize)

	// var wg sync.WaitGroup
	// wg.Add(2)

	// _ = bus.Subscribe("topic", func(v bool) {
	//     defer wg.Done()
	//     fmt.Println(v)
	// })

	// _ = bus.Subscribe("topic", func(v bool) {
	//     defer wg.Done()
	//     fmt.Println(v)
	// })

	// bus.Publish("topic", true)
	// wg.Wait()
}
