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

type Ccm struct {
	Role 		string
	CcmIP 		string
	SWVersion 	string
}


type ZoneInfo struct {
	Zid 	string
	Vip  	string
	Ccm1 	Ccm
	Ccm2 	Ccm	
}

type CardDef struct {
	Enterprise  string
	Vip		 	string
	SWVersion   string
	Name		string
	Zones		[]ZoneInfo
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
	m.Cards = []CardDef{}
	
	rslt := db.ReadDatabase[db.TBL_EnterpriseList](db.SQL_QUERIES_LOCAL["QUERY_5"].Qry)

	// Range over list of enterprise names and create a CardDef for each
	for _, result := range rslt {				
		p := CardDef{}
		p.Enterprise = result.Data[0]
		m.Cards = append(m.Cards, p)
	}	
	rslt = nil
	
	// Range over list of CardDefs and load the data for each
	for x:=0; x<len(m.Cards); x++ {

		// Check for VM, Hardware, or Mixed server types
		r := checkServerType(m.Cards[x].Enterprise)
		if r == "mixed" {
			m.Cards[x].VM = true
			m.Cards[x].Hardware = true
		} else if r == "vm" {
			m.Cards[x].VM = true
		} else {
			m.Cards[x].Hardware = true
		}

		// Load the data for each zone
		// Get a list of IP's based on the enterprise name
		s := fmt.Sprintf(db.SQL_QUERIES_LOCAL["QUERY_7"].Qry + "\"%s\"\n", m.Cards[x].Enterprise)
		rslt = db.ReadDatabase[db.TBL_ServerTypeList](s)
		fmt.Println("IP List: ", rslt)

		
	}
}

func checkServerType(ent string) string {
	vm := false
	hw := false

	s :=fmt.Sprintf(db.SQL_QUERIES_LOCAL["QUERY_6"].Qry + "\"%s\"\n", ent)
	rslt := db.ReadDatabase[db.TBL_ServerTypeList](s)
	for _, result := range rslt {
		if result.Data[0] == "VM" {
			vm = true
		} else {
			hw = true
		}
	}
	if vm && hw {
		return "mixed"
	} else if vm {
		return "vm"
	} else {
		return "hw"
	}
}

