package sidenav

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	db "dshusdock/tw_prac1/internal/services/database"
	logger "dshusdock/tw_prac1/internal/services/logging"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

type SideNavVwData struct {
	Type    string
	ID 		string
	Lbl     string
	Caret   bool
	Class   string
	SubLbl  []con.SubElement
	RepoDlg []string
	DBList  []string
	Htmx    []con.HtmxInfo
	EntList []string
}

type DSListData struct {
	Name     string
	Selected bool
}

type SideNav struct {
	App        *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       []SideNavVwData
	RepoDlg    []string
	DBList     []string
	Htmx       []con.HtmxInfo
}

var AppSideNav *SideNav

func init() {
	slog.Info("In sidenav init \n")

	pa := SIDE_NAV_BTN_LBL()
	// pb := SYS_SUB_BTN_LBL()

	AppSideNav = &SideNav{
		Id:         "sidenav",
		RenderFile: "side-nav-categories",
		ViewFlags:  []bool{true, true},
		Data: []SideNavVwData{
			{
				Type:    "caret",
				ID:	  	 "enterprise",
				Lbl:     pa.ENTERPRISE,
				Caret:   false,
				Class:   "fa-solid fa-chevron-right rotate_back",
				SubLbl:  nil,
				RepoDlg: []string{},
				DBList:  []string{},
				EntList: []string{},
				Htmx:    nil,
			},
			{
				Type:    "caret",
				ID:	  	 "swver",
				Lbl:     pa.SOFTWARE_VER,
				Caret:   false,
				Class:   "fa-solid fa-chevron-right rotate_back",
				SubLbl:  nil,
				RepoDlg: []string{},
				DBList:  []string{},
				EntList: []string{},
				Htmx:    nil,
			},
			// Next Element
		},
	}
}

func (m *SideNav) RegisterView(app config.AppConfig) *SideNav {
	fmt.Println("Registering AppSideNav...")
	AppSideNav.App = &app
	return AppSideNav
}

func (m *SideNav) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[SideNav] Process Request")
	s := d.Get("event")

	switch s {
	case con.EVENT_CLICK:
		m.processClickEvent(w, d)
	}
}

func (m *SideNav) processClickEvent(w http.ResponseWriter, d url.Values) {

	fmt.Println("\n[SideNav] ProcessClickEvent")
	lbl := d.Get("label")
	id := d.Get("view_str")

	switch d.Get("type") {
	case "caret":
		x := indexOf(lbl, m.Data)
		m.toggleCaret(x)
		m.LoadDropdownData(x)
		
		render.RenderTemplate_new(w, nil, m.Data[x], constants.RM_PARTIAL1)
	case "button":
		fmt.Printf("In the button case - %s\n", id)
		var str string
		switch id {
		case "enterprise":
			str = fmt.Sprintf("Select * from LabSystem where Enterprise = \"%s\"", lbl)
		case "swver":
			str = fmt.Sprintf("Select * from LabSystem where swVer = \"%s\"", lbl)
		}		
		
		labsystemvw.AppLSTableVW.LoadTblDataByQuery(str)
		render.RenderTemplate_new(w, nil, m.App, constants.RM_TABLE_REFRESH)
		
	case "select":
		fmt.Println("In the select case")

	default:

		//tablevw.AppTableVw.LoadTableData(lbl)
		//render.RenderMain(w, nil, m.App)
	}
}

func (m *SideNav) toggleCaret(x int) {

	if !m.Data[x].Caret {
		m.Data[x].Class = "fa fa-chevron-right rotate_fwd"
		m.Data[x].Caret = true
	} else {
		m.Data[x].Class = "fa fa-chevron-right rotate_back"
		m.Data[x].Caret = false
	}
}

func indexOf(element string, data []SideNavVwData) int {
	for k, v := range data {
		if element == v.Lbl {
			return k
		}
	}
	return -1 //not found.
}

func (m *SideNav) LoadDropdownData(x int) {
	var rslt []con.RowData

	switch x {
	case 0:
		rslt = db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_1.Qry)
	case 1:
		rslt = db.ReadDatabase[db.TBL_SWVerList](db.TBL_LAB_SYSTEM_QRY().QUERY_4.Qry)
	}

	m.Data[x].EntList = nil

	for i, result := range rslt {
		logger.Log("Result: %d %d  %s\n", x,  i, result.Data[0])
		m.Data[x].EntList = append(m.Data[x].EntList, result.Data[0])
	}
	
}
