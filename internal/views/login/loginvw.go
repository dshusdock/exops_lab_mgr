package login

import (
	"crypto/rand"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	am "dshusdock/tw_prac1/internal/services/account_mgmt"
	"dshusdock/tw_prac1/internal/services/jwtauthsvc"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/session"
	s "dshusdock/tw_prac1/internal/services/session"
	"dshusdock/tw_prac1/internal/services/token"
	"dshusdock/tw_prac1/internal/views/base"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"
	// "net/url"
)

type LoginVw struct {
	App *config.AppConfig
}

var AppLoginVw *LoginVw

func init() {
	AppLoginVw = &LoginVw{
		App: nil,
	}
	gob.Register(LoginVwData{})
	messagebus.GetBus().Subscribe("Event:ViewChange", AppLoginVw.HandleMBusRequest)

}


func (m *LoginVw) RegisterView(app *config.AppConfig) con.ViewInterface {
	log.Println("Registering AppLoginVw...")
	AppLoginVw.App = app
	return AppLoginVw
}

func (m *LoginVw) RegisterHandler() constants.ViewHandler {
	return &LoginVw{}
}

func (m *LoginVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[loginvw] - Processing request")
	CreateLoginVwData().ProcessHttpRequest(w, r)
}

func (m *LoginVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{
	CreateLoginVwData().ProcessMBusRequest(w, r)
	return nil
}

func (m *LoginVw) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[loginvw] - HandleRequest")

	var obj LoginVwData 

	if session.SessionSvc.SessionMgr.Exists(r.Context(), "loginvw") {
		obj = session.SessionSvc.SessionMgr.Pop(r.Context(), "loginvw").(LoginVwData)
	} else {
		obj = *CreateLoginVwData()	
	}

	obj.ProcessHttpRequest(w, r)	
	
	session.SessionSvc.SessionMgr.Put(r.Context(), "loginvw", obj)
	return obj
}
 

///////////////////// Login View Data //////////////////////

type LoginVwData struct {
	Base base.BaseTemplateparams
	Data any
	View int
}

func CreateLoginVwData() *LoginVwData {
	return &LoginVwData{
		Base: base.GetBaseTemplateObj(),
		Data: nil,
		View: con.RM_HOME,
	}
}

func (m *LoginVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *LoginVwData{
	fmt.Println("[loginvwData] - Processing Http request")
	var req string
	var fd url.Values

	if r.URL.Path == "/" {
		fmt.Println("Displaying login page")
		m.Base.DisplayLogin = true
		m.View = con.RM_HOME
		return m
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
		m.Base.DisplayLogin = false
		m.Base.DisplayCreateAccount = true
		m.Base.DisplayCreatAcctResponse = false
		m.View = con.RM_LOGIN

	case "create-account":
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
			return m
		}

		ai.Password = pw
		fmt.Println("Account Info: ", ai)

		// Save the account info to the database
		err = am.CreateAccount(ai)
		if err != nil {
			fmt.Println(err)
			// render.RenderTemplate_new(w, r, m, con.RM_HOME)	
			// rv.RenderViewSvc.RenderTemplate(w, r, m.Base, con.RM_HOME)
			return m
		}
				
		// If this is successful, then we will display the response
		// Eventaully I think I will want to change this approach
		m.Base.DisplayLogin = false
		m.Base.DisplayCreateAccount = false
		m.Base.DisplayCreatAcctResponse = true
		m.View = con.RM_LOGIN
	case "login":
		// Check the username and password
		if am.ValidateUser(fd.Get("username"), fd.Get("password")) {
			fmt.Println("User is valid")

			m.Base.LoggedIn = true
			m.Base.DisplayLogin = false
			m.Base.DisplayCreateAccount = false
			m.Base.DisplayCreatAcctResponse = false

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
			
			err := s.SessionSvc.SessionMgr.RenewToken(r.Context())
			// err := m.App.SessionManager.RenewToken(r.Context())
			if err != nil {
				http.Error(w, err.Error(), 500)
				return m
			}
			str, _ := GenerateRandomString(20)
			fmt.Println("Random string:", str)

			s.SessionSvc.SessionMgr.Put(r.Context(), "jwt", token)
			s.SessionSvc.SessionMgr.Put(r.Context(), "LoggedIn", true)
			s.SessionSvc.SessionMgr.Put(r.Context(), "userID", str)
			m.View = con.RM_HOME
			return m
		}
		m.Base.LoggedIn = false
		m.Base.DisplayLogin = true		
		m.View = con.RM_LOGIN
	case "logoff":
		m.Base.LoggedIn = false
		m.Base.DisplayLogin = true	
		userId := s.SessionSvc.SessionMgr.GetString(r.Context(), "userID")	
		s.SessionSvc.DeleteUserSessions(w, r, userId)
	}
	return m
}

func (m *LoginVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {}


func GenerateRandomString(length int) (string, error) {
	// Create a byte slice to store random bytes
	randomBytes := make([]byte, length)

	// Read random bytes from the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes into a base64 string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Return the random string
	return randomString[:length], nil
}





// func (m *LoginVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
// 	var req string
// 	var fd url.Values

// 	slog.Info("[loginvw] - Processing request")

// 	if r.URL.Path == "/" {
// 		fmt.Println("Displaying login page")
// 		AppLoginVw.DisplayLogin = true
// 		render.RenderTemplate_new(w, nil, AppLoginVw, con.RM_LOGIN)
// 		//rv.RenderViewPtr.UpdateView(AppLoginVw).RenderTemplate(w, r, con.RM_LOGIN)
// 		return
// 	}

// 	if r.URL.Path == "/login" {
// 		fmt.Println("Login attempt")
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		req = "login"
// 	    fd = r.PostForm
// 		slog.Info("[loginvw] - Req: ", "username", fd.Get("username"))
// 	} 

// 	if r.URL.Path == "/logoff" {
// 		slog.Info("Logoff attempt")
// 		req = "logoff"
// 	}

// 	if r.URL.Path == "/create-account" {
// 		slog.Info("Create account attempt")
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	    fd = r.PostForm
// 		req = "create-account"
// 	}

// 	if r.URL.Path == "/create-account-request" {
// 		slog.Info("Create account request attempt")
// 		req = "create-account-request"
// 	}
	
// 	switch req {
// 	case "create-account-request":
// 		obj := getLoginVwObj()
// 		obj.DisplayLogin = false
// 		obj.DisplayCreateAccount = true
// 		obj.DisplayCreatAcctResponse = false

// 		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
// 	case "create-account":
// 		obj := getLoginVwObj()
// 		// Create a new account
// 		// Need to do some validation here
// 		ai := con.AccountInfo{
// 			FirstName: fd.Get("firstname"), 
// 			LastName: fd.Get("lastname"), 
// 			Email: fd.Get("email"), 
// 			Username: fd.Get("username"), 
// 		}
// 		pw, err := token.EncryptValue(fd.Get("password"))
// 		if err != nil { 
// 			fmt.Println(err)
// 			return
// 		}

// 		ai.Password = pw

// 		fmt.Println("Account Info: ", ai)

// 		// Save the account info to the database
// 		err = am.CreateAccount(ai)
// 		if err != nil {
// 			fmt.Println(err)
// 			render.RenderTemplate_new(w, r, obj, con.RM_HOME)	
// 			return
// 		}
				
// 		// If this is successful, then we will display the response
// 		// Eventaully I think I will want to change this approach
// 		obj.DisplayLogin = false
// 		obj.DisplayCreateAccount = false
// 		obj.DisplayCreatAcctResponse = true
		
// 		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
// 	case "login":
		
// 		obj := getLoginVwObj()
// 		// Check the username and password
// 		if am.ValidateUser(fd.Get("username"), fd.Get("password")) {
// 			fmt.Println("User is valid")

// 			obj.LoggedIn = true
// 			obj.DisplayLogin = false
// 			obj.DisplayCreateAccount = false
// 			obj.DisplayCreatAcctResponse = false

// 			token, _ := jwtauthsvc.CreateToken(fd.Get("username"))
// 			http.SetCookie(w, &http.Cookie{
// 				HttpOnly: true,
// 				Expires: time.Now().Add(7 * 24 * time.Hour),
// 				SameSite: http.SameSiteLaxMode,
// 				// Uncomment below for HTTPS:
// 				Secure: true,
// 				// Must be named "jwt" or else the token cannot be 
// 				// searched for by jwtauth.Verifier.
// 				Name:  "jwt", 
// 				Value: token,
// 			})
			
// 			err := m.App.SessionManager.RenewToken(r.Context())
// 			if err != nil {
// 				http.Error(w, err.Error(), 500)
// 				return
// 			}
			
// 			m.App.SessionManager.Put(r.Context(), "jwt", token)
// 			m.App.SessionManager.Put(r.Context(), "LoggedIn", true)
			
// 			render.RenderTemplate_new(w, r, obj, con.RM_HOME)	
// 			return
// 		}
// 		obj.LoggedIn = false
// 		obj.DisplayLogin = true
		
// 		render.RenderTemplate_new(w, nil, obj, con.RM_LOGIN)
// 	case "logoff":
// 		// obj := getLoginVwObj()
// 		AppLoginVw.LoggedIn = false
// 		AppLoginVw.DisplayLogin = true
		
// 		render.RenderTemplate_new(w, nil, AppLoginVw, con.RM_LOGIN)
// 	}
// }

