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
				Lbl:     pa.ENTERPRISE,
				Caret:   true,
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

	switch d.Get("type") {
	case "caret":
		x := indexOf(lbl, m.Data)
		
		m.LoadDropdownData()

		m.toggleCaret(x)
		if m.Data[x].Caret {
			m.Data[x].Caret = false
		} else {
			m.Data[x].Caret = true
		}
		render.RenderTemplate_new(w, nil, m.App, constants.RM_SNIPPET1)
	case "button":
		fmt.Printf("In the button case - %s\n", lbl)
		str := fmt.Sprintf("Select * from LabSystem where Enterprise = \"%s\"", lbl)
		labsystemvw.AppLSTableVW.LoadTblDataByQuery(str)
		render.RenderTemplate_new(w, nil, m.App, constants.RM_TABLE_REFRESH)
		// render.RenderTemplate_new(w, nil, nil, constants.RM_SNIPPET1)

	case "select":
		fmt.Println("In the select case")

	default:

		//tablevw.AppTableVw.LoadTableData(lbl)
		//render.RenderMain(w, nil, m.App)
	}
}

func (m *SideNav) toggleCaret(x int) {

	if m.Data[x].Class == "fa fa-chevron-right rotate_back" {
		m.Data[x].Class = "fa fa-chevron-right rotate_fwd"
	} else {
		m.Data[x].Class = "fa fa-chevron-right rotate_back"
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

func (m *SideNav) LoadDropdownData() {
	rslt := db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_1.Qry)
	m.Data[0].EntList = nil

	for i, result := range rslt {
		// fmt.Printf("Result:%d  %s\n", i, result.Data[0])
		logger.Log("Result:%d  %s\n", i, result.Data[0])

		m.Data[0].EntList = append(m.Data[0].EntList, result.Data[0])

	}
	// populate the structure
	//m.Data[0].EntList = rslt[0].Data
}
