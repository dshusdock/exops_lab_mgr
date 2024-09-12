package main

import (
	"context"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/handlers"

	// "dshusdock/tw_prac1/internal/services/token"
	"dshusdock/tw_prac1/internal/services/jwtauthsvc"
	"dshusdock/tw_prac1/internal/services/unigy/unigydata"
	"dshusdock/tw_prac1/internal/services/unigy/unigystatus"
	"dshusdock/tw_prac1/internal/views/cardsvw"
	"dshusdock/tw_prac1/internal/views/headervw"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"dshusdock/tw_prac1/internal/views/layoutvw"
	"dshusdock/tw_prac1/internal/views/login"
	"dshusdock/tw_prac1/internal/views/settingsvw"
	"dshusdock/tw_prac1/internal/views/sidenav"

	// "time"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/jwtauth"
	"github.com/go-chi/jwtauth/v5"
	// "github.com/lestrrat-go/jwx/v2/jwt"
)

func initRouteHandlers() {
	// Register the views
	app.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(app)
	app.ViewCache["headervw"] = headervw.AppHeaderVw.RegisterView(app)
	app.ViewCache["settingsvw"] = settingsvw.AppSettingsVw.RegisterView(app) 
	app.ViewCache["lstablevw"] = labsystemvw.AppLSTableVW.RegisterView(app)
	app.ViewCache["sidenav"] = sidenav.AppSideNav.RegisterView(app)	
	app.ViewCache["cardsvw"] = cardsvw.AppCardsVW.RegisterView(app)
	app.ViewCache["loginvw"] = login.AppLoginVw.RegisterView(app)

	// Register the services
	app.ViewCache["unigystatus"] = unigystatus.AppStatusSvc.RegisterService(app)
	app.ViewCache["unigydata"] = unigydata.AppUnigyDataSvc.RegisterService(app)

}

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Heartbeat("/health"))
	mux.Use(middleware.Logger)
	
	// Protecting the routes
	mux.Group(func(r chi.Router) {
		// r.Use(MyMiddleware)
		
		r.Use(jwtauth.Verifier(jwtauthsvc.GetToken()))
		r.Use(jwtauth.Authenticator(jwtauthsvc.GetToken()))
		r.Get("/test", handlers.Repo.Test)

		r.Post("/logoff", handlers.Repo.Logoff)
		r.Post("/request/status", handlers.Repo.StatusInfo)
		r.Post("/upload", handlers.Repo.Upload)
		r.Post("/element/event/click", handlers.Repo.HandleClickEvents)
		r.Post("/element/event/search", handlers.Repo.HandleSearchEvents)
	})
	
	// mux.Mount("/debug", middleware.Profiler())

	mux.Get("/test2", handlers.Repo.Test2)	

	mux.Get("/", handlers.Repo.Login)
	mux.Post("/login", handlers.Repo.Login)
	
	// mux.Post("/logoff", handlers.Repo.Logoff)
	// mux.Post("/request/status", handlers.Repo.StatusInfo)
	// mux.Post("/upload", handlers.Repo.Upload)
	// mux.Post("/element/event/click", handlers.Repo.HandleClickEvents)
	// mux.Post("/element/event/search", handlers.Repo.HandleSearchEvents)

	fileServer := http.FileServer(http.Dir("./ui/html/"))
	mux.Handle("/html/*", http.StripPrefix("/html", fileServer))

	return mux
}

////////////////////////////////////////////////////////////////////////////

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Missing authorization header")
	  return
	}
	tokenString = tokenString[len("Bearer "):]
	
	err := jwtauthsvc.VerifyToken(tokenString)
	if err != nil {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Invalid token")
	  return
	}	
	fmt.Fprint(w, "Welcome to the the protected area")	
}

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

  
