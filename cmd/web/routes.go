package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/handlers"
	"dshusdock/tw_prac1/internal/services/token"
	"dshusdock/tw_prac1/internal/services/unigy/unigydata"
	"dshusdock/tw_prac1/internal/services/unigy/unigystatus"
	"dshusdock/tw_prac1/internal/views/cardsvw"
	"dshusdock/tw_prac1/internal/views/headervw"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"dshusdock/tw_prac1/internal/views/layoutvw"
	"dshusdock/tw_prac1/internal/views/settingsvw"
	"dshusdock/tw_prac1/internal/views/sidenav"

	"fmt"
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

	mux.Get("/", handlers.Repo.Login)
	mux.Get("/test", handlers.Repo.Test)
	mux.Post("/login", handlers.Repo.Home)
	mux.Post("/logoff", handlers.Repo.Logoff)
	mux.Post("/request/status", handlers.Repo.StatusInfo)
	mux.Post("/upload", handlers.Repo.Upload)
	mux.Post("/element/event/click", handlers.Repo.HandleClickEvents)
	mux.Post("/element/event/search", handlers.Repo.HandleSearchEvents)

	fileServer := http.FileServer(http.Dir("./ui/html/"))
	mux.Handle("/html/*", http.StripPrefix("/html", fileServer))

	return mux
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Missing authorization header")
	  return
	}
	tokenString = tokenString[len("Bearer "):]
	
	err := token.VerifyToken(tokenString)
	if err != nil {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Invalid token")
	  return
	}
	
	fmt.Fprint(w, "Welcome to the the protected area")
	
  }
