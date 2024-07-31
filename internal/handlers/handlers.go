package handlers

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers/upload"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
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
	// fmt.Println("Initializing handlers")
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

	messagebus.GetBus().Publish("Event:Click", w, data)
	messagebus.GetBus().Publish("Event:Change")

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
	
	render.RenderTemplate_new(w, r, m.App, con.RM_HOME)
}

func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is a test")
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
