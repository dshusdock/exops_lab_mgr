package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers"
	d "dshusdock/tw_prac1/internal/services/database"

	// "dshusdock/tw_prac1/internal/views/cardsvw"

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

	slog.Info("Starting application -", "Port", portNumber)
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
	d.ConnectLocalDB("127.0.0.1")

	// sidenav.AppSideNav.LoadDropdownData()
	//cardsvw.AppCardsVW.LoadCardDefData()
}
