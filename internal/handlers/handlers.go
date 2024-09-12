package handlers

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers/upload"
	"dshusdock/tw_prac1/internal/render"
	am "dshusdock/tw_prac1/internal/services/account_mgmt"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/jwtauthsvc"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	MBS messagebus.MessageBusSvc
}

// http.ResponseWriter, r *http.Request NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func init() {
	fmt.Println("Initializing handlers")
}	

/**
 * 	HandleClickEvents
 */
func (m *Repository) HandleClickEvents(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", con.EVENT_CLICK)
	v_id := data.Get("view_id")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return
	}

	// messagebus.GetBus().Publish("Event:Click", w, data)
	// messagebus.GetBus().Publish("Event:Change")

	// route request to appropriate handler
	ptr := m.App.ViewCache[v_id]
	ptr.ProcessRequest(w, data)

	
}

func (m *Repository) HandleSearchEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Search Events")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", con.EVENT_SEARCH)
	v_id := data.Get("view_id")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return 
	}

	fmt.Println("View ID - ", v_id)

	// messagebus.GetBus().Publish("Event:Search", w, data)

	// route request to appropriate handler
	ptr := m.App.ViewCache[v_id]
	ptr.ProcessRequest(w, data)
	
}

/**
 * 	Home is the handler for the home page
 */
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	m.App.LoggedIn = true
	render.RenderTemplate_new(w, r, m.App, con.RM_HOME)
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	d := r.PostForm
	// Check the username and password
	if am.ValidateUser(d.Get("username"), d.Get("password")) {
		fmt.Println("User is valid")
		m.App.LoggedIn = true

		token, _ := jwtauthsvc.CreateToken(d.Get("username"))

		fmt.Println("Token - ", token)

		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			Expires: time.Now().Add(7 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
			// Uncomment below for HTTPS:
			// Secure: true,
			Name:  "jwt", // Must be named "jwt" or else the token cannot be searched for by jwtauth.Verifier.
			Value: token,
		})

		render.RenderTemplate_new(w, r, m.App, con.RM_HOME)	
		return
	}
	
	render.RenderTemplate_new(w, r, m.App, con.RM_LOGIN)
}

func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n" + time.Now().String() + " - This is a test" )
	// http.Redirect(w, r, "/test2", http.StatusSeeOther)
}

func (m *Repository) Test2(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is a test2")
	
}

func (m *Repository) Logoff(w http.ResponseWriter, r *http.Request) {
	m.App.LoggedIn = false
	render.RenderTemplate_new(w, r, m.App, con.RM_HOME)
}


type Payload struct {
    Stuff string
}

func (m *Repository) StatusInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a StatusInfo")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", con.REQUEST_STATUS)
	v_id := data.Get("view_id")

	fmt.Println("View ID - ", v_id)

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		fmt.Println("No handler for route")
		return
	}
	
	// route request to appropriate handler
	ptr := m.App.ViewCache[v_id]
	ptr.ProcessRequest(w, data)
}

func testStatus(w http.ResponseWriter) {
	t := Payload{Stuff: "This is a StatusInfo"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)

	fmt.Println("This is a StatusInfo")
}

func (m *Repository) Upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is an Upload test")

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	p := upload.ProcessLabInfo(file)

	fmt.Println("Number of records - ", len(p))

}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
   
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)
	
	if u.Username == "Chek" && u.Password == "123456" {
	  tokenString, err := jwtauthsvc.CreateToken(u.Username)
	  if err != nil {
		 w.WriteHeader(http.StatusInternalServerError)
		 fmt.Errorf("No username found")
	   }
	  w.WriteHeader(http.StatusOK)
	  fmt.Fprint(w, tokenString)
	  return
	} else {
	  w.WriteHeader(http.StatusUnauthorized)
	  fmt.Fprint(w, "Invalid credentials")
	}
  }
