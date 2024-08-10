package cardsvw

import (
	"dshusdock/tw_prac1/config"
	"log/slog"

	// "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	d "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
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
	Data		[]q.DataVw1
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

	messagebus.GetBus().Subscribe("Event:ViewChange", AppCardsVW.ProcessViewChangeRequest)
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

func (m *CardsVW) ProcessViewChangeRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("[ProcessViewChangeRequest", "ID", m.Id)
	s := d.Get("label")
	slog.Info("Target - ", "Label", s)


	AppCardsVW.LoadCardData()
	render.RenderTemplate_new(w, nil, m.App, constants.RM_CARDS)
}

func (m *CardsVW) LoadCardData() {
	m.Cards = []CardDef{}
	
	rslt := d.ReadLocalDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)

	// Range over list of enterprise names and create a CardDef for each
	for _, result := range rslt {				
		p := CardDef{}
		p.Enterprise = result.Data[0]
		m.Cards = append(m.Cards, p)
	}	
	
	// Range over list of CardDefs and load the data for each
	for x:=0; x<len(m.Cards); x++ {
		fmt.Printf("----------------------Enterprise: %s ----------------------\n", m.Cards[x].Enterprise)

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
		LoadZoneData(&m.Cards[x])	
	}
}

func LoadZoneData(ptr *CardDef) {
	// Get the number of zones for the enterprise and the zone id's
		
	//  1 - Get a list of IP's based on the enterprise name
	s := fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_7"].Qry + "\"%s\"\n", ptr.Enterprise)
	rslt := d.ReadLocalDBwithType[q.TBL_ServerTypeList](s)
	count := 0
	for _, result := range rslt {
		err := d.ConnectUnigyDB(result.Data[0])
		if err != nil {
			count++
			if count > 3 {
				break
			}
			continue
		}

		// Found server to talk to 
		//  2 - Get the zone id's for the enterprise
		s := fmt.Sprintf(q.SQL_QUERIES_UNIGY["QUERY_1"].Qry )
		da := d.ReadUnigyDBwithType[q.TBL_NZData](s)
		
		for _, el := range da {
			ptr.Zones = append(ptr.Zones, 
				ZoneInfo{
					Zid: el.Data[3], 
					Vip: el.Data[2], 
					Ccm1: Ccm{CcmIP: el.Data[0]}, 
					Ccm2:Ccm{CcmIP: el.Data[0]},
				})
		}		
		d.CloseUnigyDB(result.Data[0])
	}		
}

func checkServerType(ent string) string {
	vm := false
	hw := false

	s :=fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_6"].Qry + "\"%s\"\n", ent)
	rslt := d.ReadLocalDBwithType[q.TBL_ServerTypeList](s)
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

func getZoneInfo() []ZoneInfo {

	z := []ZoneInfo{}
	return z

}

