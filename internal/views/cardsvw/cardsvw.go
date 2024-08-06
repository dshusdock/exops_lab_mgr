package cardsvw

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	db "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type CardDef struct {
	Enterprise  string
	Vip		 	string
	SWVersion   string
	Name		string
	HaPair		bool
	Active		bool
	Standby		bool
	VM     		bool
	Hardware	bool      
	Start       int
	End         int
	Query       string
	SearchInput string
	Width       []int
}

type CardsVW struct {
	App        	*config.AppConfig
	Id         	string
	RenderFile 	string
	ViewFlags  	[]bool
	Cards       []CardDef
	Data		[]db.DataVw1
	Htmx       	any
}

var AppCardsVW *CardsVW
 
func init() {

	AppCardsVW = &CardsVW{
		App:        &config.AppConfig{},
		Id:         "cardsvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Cards:      []CardDef{},
		Htmx:       nil,
	}
}

func (m *CardsVW) RegisterView(app config.AppConfig) *CardsVW {
	log.Println("Registering AppCardsVW...")
	AppCardsVW.App = &app
	return AppCardsVW
}

func (m *CardsVW) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[LSTableVW] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, m.App, con.RM_UPLOAD_MODAL)
	}
}

func (m *CardsVW) LoadCardDefData() {

	rslt := db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_5.Qry)
	for _, result := range rslt {				
		p := CardDef{}
		p.Enterprise = result.Data[0]

		m.Cards = append(m.Cards, p)
	}
	
	rslt = nil
	
	for x:=0; x<len(m.Cards); x++ {
		m.Cards[x].VM = false
		m.Cards[x].Hardware = false
		s :=fmt.Sprintf(db.TBL_LAB_SYSTEM_QRY().QUERY_6.Qry + "\"%s\"\n", m.Cards[x].Enterprise)
		rslt := db.ReadDatabase[db.TBL_ServerTypeList](s)
		for _, result := range rslt {
			if result.Data[0] == "VM" {
				m.Cards[x].VM = true
			} else {
				m.Cards[x].Hardware = true
			}
		}
	}
}

