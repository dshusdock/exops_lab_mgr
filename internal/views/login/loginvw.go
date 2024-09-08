package login

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	am "dshusdock/tw_prac1/internal/services/account_mgmt"
	"dshusdock/tw_prac1/internal/services/token"
	"fmt"
	"log"
	"net/http"
	"net/url"
)


type LoginVw struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppLoginVw *LoginVw

func init() {
	AppLoginVw = &LoginVw{
		Id:         "loginvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *LoginVw) RegisterView(app config.AppConfig) *LoginVw{
	log.Println("Registering AppLoginVw...")
	AppLoginVw.App = &app
	return AppLoginVw
}

func (m *LoginVw) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[loginvw] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)
	switch s {
	case "create-account-request":
		m.App.DisplayLogin = false
		m.App.DisplayCreateAccount = true
		m.App.DisplayCreatAcctResponse = false

		render.RenderTemplate_new(w, nil, m.App, con.RM_LOGIN)
	case "create-account":
		// Create a new account
		// Need to do some validation here
		ai := con.AccountInfo{
			FirstName: d.Get("firstname"), 
			LastName: d.Get("lastname"), 
			Email: d.Get("email"), 
			Username: d.Get("username"), 
		}
		pw, err := token.EncryptValue(d.Get("password"))
		if err != nil { 
			fmt.Println(err)
			return
		}

		ai.Password = pw

		fmt.Println("Account Info: ", ai)

		// Save the account info to the database
		am.CreateAccount(ai)

				
		// If this is successful, then we will display the response
		// Eventaully I think I will want to change this approach
		m.App.DisplayLogin = false
		m.App.DisplayCreateAccount = false
		m.App.DisplayCreatAcctResponse = true
		
		render.RenderTemplate_new(w, nil, m.App, con.RM_LOGIN)
	case "login":
		m.App.DisplayLogin = true
		m.App.DisplayCreateAccount = false
		m.App.DisplayCreatAcctResponse = false
		
		render.RenderTemplate_new(w, nil, m.App, con.RM_LOGIN)
	default:
	
		
	}
}

