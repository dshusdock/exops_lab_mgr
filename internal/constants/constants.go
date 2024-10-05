package constants

import (
	"net/http"
	// "net/url"
)

const TESTTHIS = "testthis"

const (
	FILESA = iota
	FILESB
	FILESC
)

const (
	EVENT_CLICK  = "Event_Click"
	EVENT_SEARCH = "Event_Search"
	REQUEST_STATUS = "Request_Status"
)

const (
	VW_INDEX = iota
	VW_APPHEADER
	VW_TABLE
	VW_SIDENAV
)

type EventData struct {
	Id        string
	EventType string
	Event     string
}

type HtmxInfo struct {
	Url string
}

type SubElement struct {
	Type string
	Lbl  string
}

type ViewInteface interface {
	// ProcessRequest(w http.ResponseWriter, d url.Values /*ViewInfo*/)
	ProcessRequest(w http.ResponseWriter, r *http.Request)
}



type ViewInfo struct {
	Event   int
	Type    string
	Label   string
	ViewId  string
	ViewStr string
}

type RowData struct {
	Data []string
}

// /////////////Rendered File Map///////////////
const (
	RM_HOME = iota
	RM_LOGIN
	RM_ACCOUNT_CREATE
	RM_TABLE
	RM_TABLE_REFRESH
	RM_ADD_FORM
	RM_UPLOAD_MODAL
	RM_SETTINGS_MODAL
	RM_SNIPPET1
	RM_SNIPPET2
	RM_SNIPPET3
	RM_PARTIAL1
	RM_CARDS
	RM_CARDS_MAX
	RM_CARDS_UNIGY
	RM_CARDS_SIDENAV
)

type RenderedFileMap struct {
	HOME           []string
	LOGIN		   []string
	ACCOUNT_CREATE []string
	TABLE          []string
	TABLE_REFRESH  []string
	ADD_FORM       []string
	UPLOAD_MODAL   []string
	SETTINGS_MODAL []string
	SNIPPET1       []string
	SNIPPET2       []string
	SNIPPET3       []string
	PARTIAL1       []string
	CARDS          []string
	CARDS_MAX      []string
	CARDS_UNIGY    []string
	CARDS_SIDENAV  []string
}

func RENDERED_FILE_MAP() *RenderedFileMap {
	return &RenderedFileMap{
		HOME: []string{
			"./ui/html/pages/base.tmpl.html",
			"./ui/html/pages/layout.tmpl.html",
			"./ui/html/pages/test/page1.tmpl.html",
			"./ui/html/pages/header.tmpl.html",
			"./ui/html/pages/sidenav.tmpl.html",
			"./ui/html/pages/system-list.tmpl.html",
			"./ui/html/pages/test-modal.tmpl.html",
			"./ui/html/pages/side-nav-categories.tmpl.html",
			"./ui/html/pages/lstable.tmpl.html",
			"./ui/html/pages/activity.tmpl.html",
			"./ui/html/pages/login.tmpl.html",
		},
		TABLE: []string{
			// "./ui/html/pages/layout.tmpl.html",			
			"./ui/html/pages/sidenav.tmpl.html",
			"./ui/html/pages/side-nav-categories.tmpl.html",		
			"./ui/html/pages/activity.tmpl.html",	
			"./ui/html/pages/lstable.tmpl.html",
			"./ui/html/pages/partial1.tmpl.html",
			"./ui/html/pages/cards/cards.tmpl.html",
			"./ui/html/pages/cards/unigy-cards.tmpl.html",
			"./ui/html/pages/cards/cards-snippets.tmpl.html",
		},
		TABLE_REFRESH: []string{
			"./ui/html/pages/lstable.tmpl.html",
		},
		ADD_FORM: []string{
			"./ui/html/pages/add_from.tmpl.html",
		},
		UPLOAD_MODAL: []string{
			"./ui/html/pages/test-modal.tmpl.html",
		},
		SETTINGS_MODAL: []string{
			"./ui/html/pages/settings-modal.tmpl.html",
		},
		SNIPPET1: []string{
			"./ui/html/pages/side-nav-categories.tmpl.html",
			"./ui/html/pages/partial1.tmpl.html",
		},
		SNIPPET2: []string{
			"./ui/html/pages/snippets.tmpl.html",
		},
		SNIPPET3: []string{
			"./ui/html/pages/snippets.tmpl.html",
		},
		PARTIAL1: []string{
			"./ui/html/pages/partial1.tmpl.html",
		},
		CARDS: []string{
			"./ui/html/pages/cards/cards.tmpl.html",
			"./ui/html/pages/cards/unigy-cards.tmpl.html",
			"./ui/html/pages/components/turret-card-comp.tmpl.html",
			"./ui/html/pages/cards/max-cards.tmpl.html",
			"./ui/html/pages/cards/cards-snippets.tmpl.html",
		},
		CARDS_MAX: []string{
			"./ui/html/pages/cards/cards.tmpl.html",
			"./ui/html/pages/components/turret-card-comp.tmpl.html",
			"./ui/html/pages/cards/max-cards.tmpl.html",					
		},
		CARDS_UNIGY: []string{
			"./ui/html/pages/cards/unigy-cards.tmpl.html",
		},
		CARDS_SIDENAV: []string{
			"./ui/html/pages/cards/cards-snippets.tmpl.html",
		},
		LOGIN: []string{
			"./ui/html/pages/login.tmpl.html",
		},
		ACCOUNT_CREATE: []string{
			"./ui/html/pages/login.tmpl.html",
		},
	}
}

type HeaderDef struct {
	Header string
	Width  string
}

type LocalZoneData struct {
	Id 			int
	Enterprise 	string
	Zid 		string
	Vip  		string
	Ccm1 		string
	Ccm2 		string
	Ccm1_state 	string
	Ccm2_state 	string
	Online 		bool
	Status 		string
}


type ZoneInfo struct {
	Id 			int
	Enterprise 	string
	Zid 		string
	Vip  		string
	Ccm1 		Server
	Ccm2 		Server
	Online 		bool
	Status 		string
}

type Server struct { 
	IP 			string
	SWVersion   string
	State 		string
	Active		bool
	Standby		bool
}

type AccountInfo struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  []byte
}



