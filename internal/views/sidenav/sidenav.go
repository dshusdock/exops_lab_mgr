package sidenav

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	"dshusdock/tw_prac1/internal/services/session"
	"dshusdock/tw_prac1/internal/views/base"
	"encoding/gob"

	// "dshusdock/tw_prac1/internal/views/labsystemvw"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SideNavVw struct {
	App  *config.AppConfig
}

var AppSideNavVw *SideNavVw

type SideNavBElemDetail struct {
	Type    	string
	ID 			string
	Lbl     	string
	Caret   	bool
	Class   	string
	SubLbl  	[]con.SubElement
	RepoDlg 	[]string
	DBList  	[]string
	Htmx    	[]con.HtmxInfo
	EntList 	[]string
	EntListPart []string
}

type DSListData struct {
	Name     string
	Selected bool
}

func init() {
	AppSideNavVw = &SideNavVw{
		App: nil,
	}
	gob.Register(SideNavVwData{})
}

func (m *SideNavVw) RegisterView(app *config.AppConfig) *SideNavVw {
	log.Println("Registering AppSideNav...")
	AppSideNavVw.App = app
	return AppSideNavVw
}

func (m *SideNavVw) RegisterHandler() constants.ViewHandler {
	return &SideNavVw{}
}

func (m *SideNavVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[SideNav] Process Request")
	CreateSideNavVwData().ProcessHttpRequest(w, r)
}

var GlobalObj *SideNavVwData

func (m *SideNavVw) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[SideNavVw] - HandleRequest")
	d := r.PostForm
	id := d.Get("view_id")

	var obj SideNavVwData

	if session.SessionSvc.SessionMgr.Exists(r.Context(), "sidenavvw") {
		obj = session.SessionSvc.SessionMgr.Pop(r.Context(), "sidenavvw").(SideNavVwData)
	} else {
		obj = *CreateSideNavVwData()	
	}

	if id == "sidenav" {
		obj.ProcessHttpRequest(w, r)	
	}
	session.SessionSvc.SessionMgr.Put(r.Context(), "sidenavvw", obj)

	return obj
}
 

///////////////////// SideNav View Data //////////////////////

type SideNavVwData struct {
	Base 		base.BaseTemplateparams
	SearchInput string
	Data 		[]SideNavBElemDetail
	View 		int
	TestStr 	string
}

func CreateSideNavVwData() *SideNavVwData {
	pa := SIDE_NAV_BTN_LBL()

	return &SideNavVwData{
		Base: base.GetBaseTemplateObj(),
		SearchInput: "",
		Data: []SideNavBElemDetail{
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

func (m *SideNavVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *SideNavVwData{
	d := r.PostForm
	s := d.Get("event")

	switch s {
	case con.EVENT_CLICK:
		m.processClickEvent(w, d)
	case con.EVENT_SEARCH:
		m.processSearchEvent(w, d)	
	}

	return m
}

func (m *SideNavVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {

}

func (m *SideNavVwData) processClickEvent(w http.ResponseWriter, d url.Values) {

	fmt.Println("[SideNav] ProcessClickEvent")
	lbl := d.Get("label")
	id := d.Get("view_id")

	switch d.Get("type") {
	case "caret":
		x := indexOf(lbl, m.Data)
		m.toggleCaret(x)
		m.LoadDropdownData(x)
		m.View = constants.RM_SNIPPET1		
	case "button":
		fmt.Printf("In the button case - %s - lbl - %s\n", id, lbl)
				
		// labsystemvw.AppLSTableVW.LoadTblDataByQuery(getListFromId(id, lbl))
		// labsystemvw.CreateLSTableVWData().LoadTblDataByQuery(getListFromId(id, lbl))

		render.RenderTemplate_new(w, nil, m, constants.RM_TABLE_REFRESH)
		
	case "select":
		fmt.Println("In the select case")

	default:
	}
}

func (m *SideNavVwData) processSearchEvent(w http.ResponseWriter, d url.Values) {
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

	render.RenderTemplate_new(w, nil, m, constants.RM_SNIPPET1)
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

func (m *SideNavVwData) toggleCaret(x int) {

	for count := 0; count < len(m.Data); count++ {
		if count != x {			
			// Setting to rotate_back
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

func (m *SideNavVwData) toggleCaretXXX(x int) {

	if !m.Data[x].Caret {
		m.Data[x].Class = "fa fa-chevron-right rotate_fwd"
		m.Data[x].Caret = true
	} else {
		m.Data[x].Class = "fa fa-chevron-right rotate_back"
		m.Data[x].Caret = false
	}
}

func indexOf(element string, data []SideNavBElemDetail) int {
	for k, v := range data {
		if element == v.Lbl {
			return k
		}
	}
	return -1 //not found.
}

func (m *SideNavVwData) getActiveListIdx() int {
	for k, v := range m.Data {
		if v.Caret {
			return k
		}
	}
	return -1 //not found.
}

func (m *SideNavVwData) LoadDropdownData(x int) {
	var rslt []con.RowData

	switch x {
	case 0:
		// rslt = db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_1.Qry)
		// rslt, _ = db.ReadDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_1"].Qry)
		rslt, _ = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("enterprise")
	case 1:
		// rslt = db.ReadDatabase[db.TBL_SWVerList](db.TBL_LAB_SYSTEM_QRY().QUERY_4.Qry)
		// rslt, _ = db.ReadDBwithType[q.TBL_SWVerList](q.SQL_QUERIES_LOCAL["QUERY_4"].Qry)
		rslt, _ = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("swversion")
	case 2:
		// rslt = db.ReadDatabase[db.TBL_EnterpriseList](db.TBL_LAB_SYSTEM_QRY().QUERY_5.Qry)
		// rslt, _ = db.ReadDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
		rslt, _ = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("enterprise_unigy")
	}

	m.Data[x].EntList = nil

	for _, result := range rslt {
		// logger.Log("Result: %d %d  %s\n", x,  i, result.Data[0])
		m.Data[x].EntList = append(m.Data[x].EntList, result.Data[0])
	}
	m.Data[x].EntListPart = m.Data[x].EntList
	
}
