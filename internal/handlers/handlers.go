package handlers

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

type LabDataRow struct {
	Cab string
	U string
	ISO string
	Name string
	SerialNbr string
	IP string
	VIP string
	IdracIP string
	SWVer string
	ServerType string
	Enterprise string
	Role string
	Comments string
	VMLabServerHostIP string
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

/**
 * 	HandleClickEvents
 */
func (m *Repository) HandleClickEvents(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", constants.EVENT_CLICK)
	v_id := data.Get("view_id")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return
	}

	// route request to appropriate handler
	ptr := m.App.ViewCache[v_id]
	ptr.ProcessRequest(w, data)
}

/**
 * 	Home is the handler for the home page
 */
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate_new(w, r, m.App, constants.RM_HOME)
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

	p := processLabInfo(file)

	fmt.Println("Number of records - ", len(p))
	
}

func processLabInfo(f multipart.File) ([]LabDataRow){
	var result []LabDataRow
	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var x = LabDataRow{}
		for i := 0; i < len(record); i++ {
			switch(i) {
			case 0:
				x.Cab = record[i]
			case 1:
				x.U = record[i]
			case 2:
				x.ISO = record[i]
			case 3:
				x.Name = record[i]
			case 4:
				x.SerialNbr = record[i]
			case 5:
				x.IP = record[i]
			case 6:
				x.VIP = record[i]
			case 7:
				x.IdracIP = record[i]
			case 8:
				x.SWVer = record[i]
			case 9:
				x.ServerType= record[i]
			case 10:
				x.Enterprise = record[i]
			case 11:
				x.Role = record[i]
			case 12:
				x.Comments = record[i]
			case 13:
				x.VMLabServerHostIP = record[i]				
			}
		}
		result = append(result, x)
	}
	return result
}
