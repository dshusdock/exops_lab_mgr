package cardsvw

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	d "dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

type CardDef struct {
	Enterprise  string
	Vip		 	string
	Name		string
	Zones		[]con.ZoneInfo
	VM     		bool
	Hardware	bool      
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
	var fileMap int16
	fmt.Println("[AppCardsVW] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		
	case "Max":
		fileMap = con.RM_CARDS_MAX
	case "Unigy":
		fileMap = con.RM_CARDS_UNIGY		
	}
	render.RenderTemplate_new(w, nil, m.App, fileMap)
}

func (m *CardsVW) ProcessViewChangeRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("[ProcessViewChangeRequest", "ID", m.Id)
	s := d.Get("label")
	slog.Info("Target - ", "Label", s)
	
	m.App.MainTable = false
	m.App.Cards = true

	AppCardsVW.LoadCardData()
	render.RenderTemplate_new(w, nil, m.App, con.RM_CARDS)
}

func (m *CardsVW) LoadCardData() error{
	slog.Info("In LoadCardData...")
	m.Cards = []CardDef{}
	
	// rslt, err := d.ReadDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
	rslt, _ := dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("enterprise_unigy")

	for _, result := range rslt {				
		p := CardDef{}
		p.Enterprise = result.Data[0]
		m.Cards = append(m.Cards, p)
	}	
	
	// Range over list of CardDefs and load the data for each
	for x:=0; x<len(m.Cards); x++ {
		// Check for VM, Hardware, or Mixed server types
		r, err := checkServerType(m.Cards[x].Enterprise)
		if err != nil {
			fmt.Println("Error in LoadCardData: ", err)
			return err
		}
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
	return nil
}

func LoadZoneData(ptr *CardDef) error{

	// s :=fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_9"].Qry + "\"%s\"", ptr.Enterprise)
	// rslt, err := d.ReadDBwithType[con.LocalZoneData](s)

	rslt, _ := dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetView(dbdata.VIEW_6, ptr.Enterprise)

	for _, result := range rslt {
		z := con.ZoneInfo{}
		z.Id, _ = strconv.Atoi(result.Data[0])
		z.Enterprise = result.Data[1]
		z.Zid = result.Data[2]
		z.Vip = result.Data[3]
		z.Ccm1 = con.Server{
			IP: result.Data[4],
			SWVersion: "" ,
			State: result.Data[6],
			Active: false,
			Standby: false,
		}
		z.Ccm2 = con.Server{
			IP: result.Data[5],
			SWVersion: "" ,
			State: result.Data[7],
			Active: false,
			Standby: false,
		}
		z.Online, _ = strconv.ParseBool(result.Data[8]) 
		z.Status = result.Data[9]
		ptr.Zones = append(ptr.Zones, z)
	}
	return nil
}

func checkServerType(ent string) (string, error){
	vm := false
	hw := false

	s :=fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_6"].Qry + "\"%s\"\n", ent)
	rslt, err := d.ReadDBwithType[q.TBL_ServerTypeList](s)
	if err != nil {
		fmt.Println("Error in checkServerType: ", err)
		return "", err
	}

	for _, result := range rslt {
		if result.Data[0] == "VM" {
			vm = true
		} else {
			hw = true
		}
	}
	if vm && hw {
		return "mixed", nil
	} else if vm {
		return "vm", nil
	} else {
		return "hw", nil
	}
}

