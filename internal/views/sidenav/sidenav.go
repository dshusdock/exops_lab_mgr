package sidenav

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	db "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	EntListPart []string
}

type DSListData struct {
	Name     string
	Selected bool
}

type SideNav struct {
	App        	*config.AppConfig
	Id         	string
	RenderFile 	string
	ViewFlags  	[]bool
	SearchInput string
	Data       	[]SideNavVwData
	RepoDlg    	[]string
	DBList     	[]string
	Htmx       	[]con.HtmxInfo
}

var AppSideNav *SideNav

func init() {
	pa := SIDE_NAV_BTN_LBL()
	// pb := SYS_SUB_BTN_LBL()

	AppSideNav = &SideNav{
		Id:         "sidenav",
		RenderFile: "side-nav-categories",
		ViewFlags:  []bool{true, true},
		SearchInput: "",
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
				EntListPart: []string{},
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
				EntListPart: []string{},
				Htmx:    nil,
			},
			{
				Type:    "caret",
				ID:	  	 "Unigy",
				Lbl:     pa.UNIGY,
				Caret:   false,
				Class:   "fa-solid fa-chevron-right rotate_back",
				SubLbl:  nil,
				RepoDlg: []string{},
				DBList:  []string{},
				EntList: []string{},
				EntListPart: []string{},
				Htmx:    nil,
			},
		},
	}
}

func (m *SideNav) RegisterView(app config.AppConfig) *SideNav {
	log.Println("Registering AppSideNav...")
	AppSideNav.App = &app
	return AppSideNav
}

func (m *SideNav) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[SideNav] Process Request")
	s := d.Get("event")

	switch s {
	case con.EVENT_CLICK:
		m.processClickEvent(w, d)
	case con.EVENT_SEARCH:
		m.processSearchEvent(w, d)	
	}
}

func (m *SideNav) processClickEvent(w http.ResponseWriter, d url.Values) {

	fmt.Println("[SideNav] ProcessClickEvent")
	lbl := d.Get("label")
	id := d.Get("view_str")

	switch d.Get("type") {
	case "caret":
		x := indexOf(lbl, m.Data)
		m.toggleCaret(x)
		m.LoadDropdownData(x)
		
		render.RenderTemplate_new(w, nil, m.App, constants.RM_SNIPPET1)
	case "button":
		fmt.Printf("In the button case - %s - lbl - %s\n", id, lbl)
				
		labsystemvw.AppLSTableVW.LoadTblDataByQuery(getListFromId(id, lbl))
		render.RenderTemplate_new(w, nil, m.App, constants.RM_TABLE_REFRESH)
		
	case "select":
		fmt.Println("In the select case")

	default:

		//tablevw.AppTableVw.LoadTableData(lbl)
		//render.RenderMain(w, nil, m.App)
	}
}

func (m *SideNav) processSearchEvent(w http.ResponseWriter, d url.Values) {
	var rd []string
	key := d.Get("search")
	lbl := d.Get("label")
	idx := m.getActiveListIdx()
	m.SearchInput = key

	if idx < 0 { return }

	fmt.Printf("Search: %s  Label: %s  Index: %d\n", key, lbl, idx)

	if key == "" {
		fmt.Println("Key is null")
		rd = m.Data[idx].EntList
	} else {
		fmt.Println("Key is not null")
		for i := 0; i < len(m.Data[idx].EntList); i++ {
			str := m.Data[idx].EntList[i]
			if strings.Contains(str, key) {				
				rd = append(rd, str)
			}
		}
	}
	fmt.Printf("Result: %v\n", rd)
	m.Data[idx].EntListPart = rd

	render.RenderTemplate_new(w, nil, m.App, constants.RM_SNIPPET1)
}

func getListFromId(id string, lbl string) string {
	var str string
	part := "Select * from LabSystem where "

	switch id {
	case "enterprise":
		str = fmt.Sprintf(part + "Enterprise = \"%s\"", lbl)
	case "swver":
		str = fmt.Sprintf(part + "swVer = \"%s\"", lbl)
	case "Unigy":
		str = fmt.Sprintf(part + "Enterprise = \"%s\"", lbl)
	}

	return str
}

func (m *SideNav) toggleCaret(x int) {

	for count := 0; count < len(m.Data); count++ {
		fmt.Printf("Count: %d  x: %d\n", count, x)
		if count != x {			
			fmt.Print("Setting to rotate_back\n")
			m.Data[count].Class = "fa fa-chevron-right rotate_back"
			m.Data[count].Caret = false			
		} else {

			if !m.Data[count].Caret {
				m.Data[count].Class = "fa fa-chevron-right rotate_fwd"
				m.Data[count].Caret = true
			} else {
				m.Data[count].Class = "fa fa-chevron-right rotate_back"
				m.Data[count].Caret = false
			}	
		}
	}
}

func (m *SideNav) toggleCaretXXX(x int) {

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

func (m *SideNav) getActiveListIdx() int {
	for k, v := range m.Data {
		if v.Caret {
			return k
		}
	}
	return -1 //not found.
}

func (m *SideNav) LoadDropdownData(x int) {
	var rslt []con.RowData

	switch x {
	case 0:
		// rslt = db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_1.Qry)
		rslt = db.ReadLocalDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_1"].Qry)
	case 1:
		// rslt = db.ReadDatabase[db.TBL_SWVerList](db.TBL_LAB_SYSTEM_QRY().QUERY_4.Qry)
		rslt = db.ReadLocalDBwithType[q.TBL_SWVerList](q.SQL_QUERIES_LOCAL["QUERY_4"].Qry)
	case 2:
		// rslt = db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_5.Qry)
		rslt = db.ReadLocalDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
	}

	m.Data[x].EntList = nil

	for _, result := range rslt {
		// logger.Log("Result: %d %d  %s\n", x,  i, result.Data[0])
		m.Data[x].EntList = append(m.Data[x].EntList, result.Data[0])
	}
	m.Data[x].EntListPart = m.Data[x].EntList
	
}
