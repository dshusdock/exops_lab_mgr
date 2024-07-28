package constants

import (
	"net/http"
	"net/url"
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
	ProcessRequest(w http.ResponseWriter, d url.Values /*ViewInfo*/)
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
	RM_TABLE
	RM_TABLE_REFRESH
	RM_ADD_FORM
	RM_UPLOAD_MODAL
	RM_SETTINGS_MODAL
	RM_SNIPPET1
)

type RenderedFileMap struct {
	HOME           []string
	TABLE          []string
	TABLE_REFRESH  []string
	ADD_FORM       []string
	UPLOAD_MODAL   []string
	SETTINGS_MODAL []string
	SNIPPET1       []string
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
		},
		TABLE: []string{
			// "./ui/html/pages/layout.tmpl.html",			
			"./ui/html/pages/sidenav.tmpl.html",
			"./ui/html/pages/side-nav-categories.tmpl.html",		
			"./ui/html/pages/activity.tmpl.html",	
			"./ui/html/pages/lstable.tmpl.html",
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
		},
	}
}

type HeaderDef struct {
	Header string
	Width  string
}
