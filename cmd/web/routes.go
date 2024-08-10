package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/handlers"
	"dshusdock/tw_prac1/internal/services/unigy/unigydata"
	"dshusdock/tw_prac1/internal/services/unigy/unigystatus"
	"dshusdock/tw_prac1/internal/views/cardsvw"
	"dshusdock/tw_prac1/internal/views/headervw"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"dshusdock/tw_prac1/internal/views/layoutvw"
	"dshusdock/tw_prac1/internal/views/settingsvw"
	"dshusdock/tw_prac1/internal/views/sidenav"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRouteHandlers() {
	// Register the views
	app.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(app)
	app.ViewCache["headervw"] = headervw.AppHeaderVw.RegisterView(app)
	app.ViewCache["settingsvw"] = settingsvw.AppSettingsVw.RegisterView(app) 
	app.ViewCache["lstablevw"] = labsystemvw.AppLSTableVW.RegisterView(app)
	app.ViewCache["sidenav"] = sidenav.AppSideNav.RegisterView(app)	
	app.ViewCache["cardsvw"] = cardsvw.AppCardsVW.RegisterView(app)

	// Register the services
	app.ViewCache["unigystatus"] = unigystatus.AppStatusSvc.RegisterService(app)
	app.ViewCache["unigydata"] = unigydata.AppUnigyDataSvc.RegisterService(app)

}

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/test", handlers.Repo.Test)
	mux.Post("/request/status", handlers.Repo.StatusInfo)
	mux.Post("/upload", handlers.Repo.Upload)
	mux.Post("/element/event/click", handlers.Repo.HandleClickEvents)
	mux.Post("/element/event/search", handlers.Repo.HandleSearchEvents)

	fileServer := http.FileServer(http.Dir("./ui/html/"))
	mux.Handle("/html/*", http.StripPrefix("/html", fileServer))

	return mux
}
