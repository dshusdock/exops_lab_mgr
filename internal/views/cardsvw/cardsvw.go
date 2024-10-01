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
	SwVer 		[]con.RowData
	Zones		[]con.ZoneInfo
	VM     		bool
	Hardware	bool   
	Display		bool   
}

type TurretDef struct {
	Enterprise  string
	IP		 	string
	ParentZone	string
	Label		string
}

type SidePanelBtn struct {
	Label string
}

const ( 
	ENTERPRISE = iota
	SWVERSION
)

type CardsVW struct {
	App        		*config.AppConfig
	Id         		string
	RenderFile 		string
	ViewFlags  		[]bool
	Cards       	[]CardDef
	Turret			[]TurretDef
	Htmx       		any
	SidePanelDef 	[]SidePanelBtn
	Radio1			bool
}

var AppCardsVW *CardsVW
 
func init() {

	AppCardsVW = &CardsVW{
		App:        &config.AppConfig{},
		Id:         "cardsvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Cards:      []CardDef{},
		Turret:		 []TurretDef{},
		Htmx:       nil,
		SidePanelDef: []SidePanelBtn{},
		Radio1:     true,
	}

	messagebus.GetBus().Subscribe("Event:ViewChange", AppCardsVW.ProcessMBusRequest)
}

func (m *CardsVW) RegisterView(app config.AppConfig) *CardsVW {
	log.Println("Registering AppCardsVW...")
	AppCardsVW.App = &app
	return AppCardsVW
}

func (m *CardsVW) ProcessRequest(w http.ResponseWriter, d url.Values) {
	var fileMap int16
	fmt.Println("[AppCardsVW] - Processing request")
	lbl := d.Get("label")
	_type := d.Get("type")
	selector := d.Get("view_str")
	fmt.Println("selector: ", selector)

	if selector == "device-selector" {
		switch _type {
		case "button":
			switch lbl {
			case "upload":
			case "Max":
				m.LoadTurretData("max")
				fileMap = con.RM_CARDS_MAX
			case "Unigy":
				AppCardsVW.LoadCardData()
				fileMap = con.RM_CARDS_UNIGY	
			case "Touch":
				fileMap = con.RM_CARDS_MAX
				m.LoadTurretData("mercury")		
			}
			render.RenderTemplate_new(w, nil, m.App, fileMap)
		}
	} else {
		switch _type {
		case "button":
			m.FilterView(lbl)
			render.RenderTemplate_new(w, nil, m.App, con.RM_CARDS_UNIGY)
		case "radio":
			fmt.Println("Radio button selected")
			fileMap = con.RM_CARDS_SIDENAV
			if lbl == "radio1" {
				AppCardsVW.UpdateSidePanel(ENTERPRISE)
				m.Radio1 = true
			} else {
				AppCardsVW.UpdateSidePanel(SWVERSION)		
				m.Radio1 = false	
			}
			render.RenderTemplate_new(w, nil, m.App, fileMap)
		}
	}
}

func (m *CardsVW) ProcessMBusRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("[ProcessViewChangeRequest", "ID", m.Id)
	s := d.Get("label")
	slog.Info("Target - ", "Label", s)
	
	m.App.MainTable = false
	m.App.Cards = true

	AppCardsVW.LoadCardData()
	AppCardsVW.UpdateSidePanel(ENTERPRISE)
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
		p.Display = true
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

		m.Cards[x].SwVer, _ = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetView(dbdata.VIEW_8, m.Cards[x].Enterprise)

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

func (m *CardsVW) LoadTurretData(t string) error{ 
	slog.Info("In LoadTurretData...")
	m.Turret = []TurretDef{}
	
	rslt, _ := dbdata.GetDBAccess(dbdata.DEVICE).GetAll()

	for _, result := range rslt {	
		if result.Data[2] == t {			
			p := TurretDef{}
			p.Enterprise = result.Data[1]
			p.IP = result.Data[5]
			p.ParentZone = result.Data[7]
			p.Label = t
			m.Turret = append(m.Turret, p)
		}
	}	
	return nil
}

func (m *CardsVW) UpdateSidePanel(viewType int) {
	m.SidePanelDef = []SidePanelBtn{}
	if viewType == ENTERPRISE {
		rsltb, _ := dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("enterprise_unigy")
		for _, result := range rsltb {
			p := SidePanelBtn{}
			p.Label = result.Data[0]
			m.SidePanelDef = append(m.SidePanelDef, p)
		}
	} else {
		rslta, _ := dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("swversion_unigy")
		for _, result := range rslta {
			p := SidePanelBtn{}
			p.Label = result.Data[0]
			m.SidePanelDef = append(m.SidePanelDef, p)
		}
	}

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

func (m *CardsVW) FilterView(lbl string) {
	fmt.Println("FilterView: ", lbl)

	if m.Radio1 {
		for i:=0; i<len(m.Cards); i++ {
			if m.Cards[i].Enterprise == lbl {				
				m.Cards[i].Display = true
			} else {
				m.Cards[i].Display = false
			}
		}
	} else {
		for i:=0; i<len(m.Cards); i++ {
			if m.Cards[i].SwVer[0].Data[0] == lbl {				
				m.Cards[i].Display = true
			} else {
				m.Cards[i].Display = false
			}
		}
	}
}