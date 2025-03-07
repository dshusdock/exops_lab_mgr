package main

import (
	"context"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/handlers"
	"dshusdock/tw_prac1/internal/services/jwtauthsvc"
	// "dshusdock/tw_prac1/internal/services/unigy/unigydata"
	// "dshusdock/tw_prac1/internal/services/unigy/unigystatus"
	// "dshusdock/tw_prac1/internal/views/base"
	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/headervw"
	// "dshusdock/tw_prac1/internal/views/labsystemvw"
	// "dshusdock/tw_prac1/internal/views/layoutvw"
	// "dshusdock/tw_prac1/internal/views/login"
	// "dshusdock/tw_prac1/internal/views/settingsvw"
	// "dshusdock/tw_prac1/internal/views/sidenav"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Heartbeat("/health"))
	mux.Use(middleware.Logger)
	
	// Protecting the routes
	mux.Group(func(r chi.Router) {
		// r.Use(MyMiddleware)		
		r.Use(jwtauth.Verifier(jwtauthsvc.GetToken()))
		r.Use(jwtauth.Authenticator(jwtauthsvc.GetToken()))	
			
		r.Post("/logoff", handlers.Repo.Logoff)		
		r.Post("/request/status", handlers.Repo.StatusInfo)
		r.Post("/upload", handlers.Repo.Upload)
		r.Post("/element/event/click", handlers.Repo.HandleClickEvents)
		r.Post("/element/event/search", handlers.Repo.HandleSearchEvents)

		r.Get("/test", handlers.Repo.Test)
	})
	// Unprotected routes
	mux.Get("/", handlers.Repo.Login)
	mux.Post("/login", handlers.Repo.Login)
	mux.Post("/create-account-request", handlers.Repo.CreateAccount)
	mux.Post("/create-account", handlers.Repo.CreateAccount)

	mux.Get("/test2", handlers.Repo.Test2)	

	fileServer := http.FileServer(http.Dir("./ui/html/"))
	mux.Handle("/html/*", http.StripPrefix("/html", fileServer))

	return mux
}

////////////////////////////////////////////////////////////////////////////

// Trying some of chi's middlewares
// HTTP middleware setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  // create new context from `r` request context, and assign key `"user"`
	  // to value of `"123"`
	  ctx := context.WithValue(r.Context(), "user", "123")
  
	  // call the next handler in the chain, passing the response writer and
	  // the updated request object with the new context value.
	  //
	  // note: context.Context values are nested, so any previously set
	  // values will be accessible as well, and the new `"user"` key
	  // will be accessible from this point forward.
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
  }

  func MyHandler(w http.ResponseWriter, r *http.Request) {
    // here we read from the request context and fetch out `"user"` key set in
    // the MyMiddleware example above.
    user := r.Context().Value("user").(string)

    // respond to the client
    w.Write([]byte(fmt.Sprintf("hi %s", user)))
}

  
