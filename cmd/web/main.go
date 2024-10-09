package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers"
	renderview "dshusdock/tw_prac1/internal/services/renderview"
	"dshusdock/tw_prac1/internal/services/session"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8084"
const secPortNumber = ":8443"

var app config.AppConfig

func main() {
	app = config.AppConfig{}

	app.InProduction = false
	app.SideNav = false
	app.MainTable = false
	app.ViewCache = make(map[string]constants.ViewInterface)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render := renderview.NewRenderViewSvc(&app)
	renderview.MapRenderViewSvc(render)
	renderview.InitRouteHandlers()
	
	// Session Manager
	app.SessionManager = scs.New()
	app.SessionManager.Lifetime = 3 * time.Hour
	app.SessionManager.IdleTimeout = 20 * time.Minute
	app.SessionManager.Cookie.Name = "session_id"
	app.SessionManager.Cookie.Domain = "10.205.185.154"
	app.SessionManager.Cookie.HttpOnly = true
	// app.SessionManager.Cookie.Path = "/exops/"
	app.SessionManager.Cookie.Persist = true
	// app.SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	app.SessionManager.Cookie.Secure = true

	session.SessionSvc.RegisterSessionManager(app.SessionManager)

	// Logging - Info by default
	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)

	// slog.Info("Starting application -", "Port", portNumber)
	slog.Info("Starting application -", "Port", secPortNumber)
	srv := &http.Server{
		// Addr:    portNumber,
		Addr:    secPortNumber,
		Handler: app.SessionManager.LoadAndSave(routes(&app)),
	}

	// err := srv.ListenAndServe()
	err := srv.ListenAndServeTLS("dev_cert.crt", "dev_key.key")
	if err != nil {
		log.Fatal(err)
	}
}
