package render

import (
	"dshusdock/tw_prac1/internal/constants"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var files = []string{
	"./ui/html/pages/base.tmpl.html",
	"./ui/html/pages/layout.tmpl.html",
	"./ui/html/pages/header.tmpl.html",
	"./ui/html/pages/test/page1.tmpl.html",
	"./ui/html/pages/sidenav.tmpl.html",
	"./ui/html/pages/system-list.tmpl.html",
	"./ui/html/pages/test-modal.tmpl.html",
}

type Payload struct {
    Server string
}

func JSONResponse(w http.ResponseWriter, data string) {
	t := Payload{Server: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, d any) {
	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "base", d)
}

func RenderModal(w http.ResponseWriter, r *http.Request, d any) {
	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "test-modal", d)
}

// RenderTemplate renders a template
func RenderTemplate_new(w http.ResponseWriter, r *http.Request, data any, _type int16) {
	ptr := constants.RENDERED_FILE_MAP()
	fmt.Println("Type: ", _type)

	var tmplFiles []string
	var tmplName string

	switch _type {
	case constants.RM_HOME:
		tmplFiles = ptr.HOME
		tmplName = "base"
	case constants.RM_LOGIN:
		tmplFiles = ptr.HOME
		tmplName = "base"
	case constants.RM_ADD_FORM:
		tmplFiles = ptr.ADD_FORM
		tmplName = "base"
	case constants.RM_UPLOAD_MODAL:
		tmplFiles = ptr.UPLOAD_MODAL
		tmplName = "test-modal"
	case constants.RM_SETTINGS_MODAL:
		tmplFiles = ptr.SETTINGS_MODAL
		tmplName = "settings-modal"
	case constants.RM_TABLE:
		tmplFiles = ptr.TABLE
		tmplName = "activity"
	case constants.RM_TABLE_REFRESH:
		tmplFiles = ptr.TABLE
		tmplName = "lstable"
	case constants.RM_SNIPPET1:
		tmplFiles = ptr.SNIPPET1
		tmplName = "side-nav-categories"
	case constants.RM_SNIPPET3:
		tmplFiles = ptr.SNIPPET3
		tmplName = "snippets3"
	case constants.RM_PARTIAL1:
		tmplFiles = ptr.PARTIAL1
		tmplName = "partial1"
	case constants.RM_CARDS:
		tmplFiles = ptr.CARDS
		tmplName = "cardsvw"
	case constants.RM_CARDS_MAX:
		tmplFiles = ptr.CARDS
		tmplName = "max_cardsvw"
	case constants.RM_CARDS_UNIGY:
		tmplFiles = ptr.CARDS
		tmplName = "tmplt_unigy-cards-view"
	case constants.RM_ACCOUNT_CREATE:
		tmplFiles = ptr.ACCOUNT_CREATE
		tmplName = "acct-create-response"
	default:
		tmplFiles = ptr.HOME
		tmplName = "base"
	}
	template.Must(template.ParseFiles(tmplFiles...)).ExecuteTemplate(w, tmplName, data)
}
