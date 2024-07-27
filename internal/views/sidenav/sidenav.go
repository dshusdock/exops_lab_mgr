package sidenav

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	db "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"log"
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
				Class:   "bi-caret-right",
				SubLbl:  nil,
				RepoDlg: []string{"border", "IP Address"},
				DBList:  []string{},
				EntList: []string{"Item1", "Item2", "Item3"},
				Htmx:    nil,
			},
			{
				Type:    "caret",
				Lbl:     pa.USER,
				Caret:   true,
				Class:   "bi-caret-right",
				SubLbl:  nil,
				RepoDlg: []string{"border", "IP Address"},
				DBList:  []string{},
				Htmx:    nil,
				EntList: []string{"Item4", "Item5", "Item6"},
			},
			// Next Element
		},
	}
}

func (m *SideNav) RegisterView(app config.AppConfig) *SideNav {
	log.Println("Registering AppSideNav...")
	AppSideNav.App = &app
	return AppSideNav
}

func (m *SideNav) ProcessRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("[SideNav] Entering Process Request")
	s := d.Get("event")

	switch s {
	case con.EVENT_CLICK:
		m.processClickEvent(w, d)
	}
}

func (m *SideNav) processClickEvent(w http.ResponseWriter, d url.Values) {

	fmt.Println("[SideNav] In processClickEvent")
	lbl := d.Get("label")

	switch d.Get("type") {
	case "caret":
		x := indexOf(lbl, m.Data)
		m.toggleCaret(x)
		if m.Data[x].Caret {
			m.Data[x].Caret = false
		} else {
			m.Data[x].Caret = true
		}
		//render.RenderSideNav(w, nil, m.App)
	case "button":
		fmt.Println("In the button case")

	case "select":
		fmt.Println("In the select case")

	default:

		//tablevw.AppTableVw.DisplaySQLTable(lbl)
		//render.RenderMain(w, nil, m.App)
	}

	fmt.Println("HERE")
}

func (m *SideNav) toggleCaret(x int) {

	if m.Data[x].Class == "bi-caret-down" {
		m.Data[x].Class = "bi-caret-right"
	} else {
		m.Data[x].Class = "bi-caret-down"
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

func (m *SideNav) InitDropdownData() {
	rslt := db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_1.Qry)

	for _, result := range rslt {
		fmt.Printf("Result: %+v\n", result)

	}

	//m.Data[0].EntList = rslt[0].Data
}
