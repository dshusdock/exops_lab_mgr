package login

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	am "dshusdock/tw_prac1/internal/services/account_mgmt"
	"dshusdock/tw_prac1/internal/services/jwtauthsvc"
	"dshusdock/tw_prac1/internal/services/token"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"
	// "net/url"
)


type LoginVw struct {
	App 						*config.AppConfig
	Id         					string
	LoggedIn 					bool
	DisplayLogin  				bool
	DisplayCreateAccount 		bool
	DisplayCreatAcctResponse 	bool
	SideNav	      				bool
	MainTable	  				bool
	Cards		  				bool
}

func getLoginVwObj() LoginVw {
	return LoginVw{
		Id:         "loginvw",
		LoggedIn: false,
		DisplayLogin: true,
		DisplayCreateAccount: false,
		DisplayCreatAcctResponse: false,
		SideNav: false,
		MainTable: false,
		Cards: false,
	}
}

var AppLoginVw *LoginVw

func init() {
	obj := getLoginVwObj()
	AppLoginVw = &obj
}

func (m *LoginVw) RegisterView(app config.AppConfig) *LoginVw{
	log.Println("Registering AppLoginVw...")
	AppLoginVw.App = &app
	return AppLoginVw
}

func (m *LoginVw) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	var req string
	var fd url.Values

	slog.Info("[loginvw] - Processing request")

	if r.URL.Path == "/" {
		fmt.Println("Displaying login page")
		AppLoginVw.DisplayLogin = true
		render.RenderTemplate_new(w, nil,AppLoginVw, con.RM_LOGIN)
		return
	}

	if r.URL.Path == "/login" {
		fmt.Println("Login attempt")
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		req = "login"

	    fd = r.PostForm
		slog.Info("[loginvw] - Req: ", "username", fd.Get("username"))
	} 

	if r.URL.Path == "/logoff" {
		slog.Info("Logoff attempt")
		req = "logoff"
	}

	if r.URL.Path == "/create-account" {
		slog.Info("Create account attempt")
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
	    fd = r.PostForm
		req = "create-account"
	}

	if r.URL.Path == "/create-account-request" {
		slog.Info("Create account request attempt")
		req = "create-account-request"
	}
	
	switch req {
	case "create-account-request":
		obj := getLoginVwObj()
		obj.DisplayLogin = false
		obj.DisplayCreateAccount = true
		obj.DisplayCreatAcctResponse = false

		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
	case "create-account":
		obj := getLoginVwObj()
		// Create a new account
		// Need to do some validation here
		ai := con.AccountInfo{
			FirstName: fd.Get("firstname"), 
			LastName: fd.Get("lastname"), 
			Email: fd.Get("email"), 
			Username: fd.Get("username"), 
		}
		pw, err := token.EncryptValue(fd.Get("password"))
		if err != nil { 
			fmt.Println(err)
			return
		}

		ai.Password = pw

		fmt.Println("Account Info: ", ai)

		// Save the account info to the database
		err = am.CreateAccount(ai)
		if err != nil {
			fmt.Println(err)
			render.RenderTemplate_new(w, r, obj, con.RM_HOME)	
			return
		}
				
		// If this is successful, then we will display the response
		// Eventaully I think I will want to change this approach
		obj.DisplayLogin = false
		obj.DisplayCreateAccount = false
		obj.DisplayCreatAcctResponse = true
		
		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
	case "login":
		
		obj := getLoginVwObj()
		// Check the username and password
		if am.ValidateUser(fd.Get("username"), fd.Get("password")) {
			fmt.Println("User is valid")

			obj.LoggedIn = true
			obj.DisplayLogin = false
			obj.DisplayCreateAccount = false
			obj.DisplayCreatAcctResponse = false

			token, _ := jwtauthsvc.CreateToken(fd.Get("username"))
			http.SetCookie(w, &http.Cookie{
				HttpOnly: true,
				Expires: time.Now().Add(7 * 24 * time.Hour),
				SameSite: http.SameSiteLaxMode,
				// Uncomment below for HTTPS:
				Secure: true,
				// Must be named "jwt" or else the token cannot be 
				// searched for by jwtauth.Verifier.
				Name:  "jwt", 
				Value: token,
			})
			
			err := m.App.SessionManager.RenewToken(r.Context())
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			m.App.SessionManager.Put(r.Context(), "jwt", token)
			m.App.SessionManager.Put(r.Context(), "LoggedIn", true)
			
			render.RenderTemplate_new(w, r, obj, con.RM_HOME)	
			return
		}
		obj.LoggedIn = false
		obj.DisplayLogin = true
		
		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
	case "logoff":
		// obj := getLoginVwObj()
		AppLoginVw.LoggedIn = false
		AppLoginVw.DisplayLogin = true
		
		render.RenderTemplate_new(w, nil, AppLoginVw, con.RM_LOGIN)

	
		
	}
}

