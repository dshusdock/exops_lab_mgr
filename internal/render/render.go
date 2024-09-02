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
	switch _type {
	case constants.RM_HOME:
		template.Must(template.ParseFiles(ptr.HOME...)).ExecuteTemplate(w, "base", data)
	case constants.RM_LOGIN:
		template.Must(template.ParseFiles(ptr.HOME...)).ExecuteTemplate(w, "base", data)
	case constants.RM_ADD_FORM:
		template.Must(template.ParseFiles(ptr.ADD_FORM...)).ExecuteTemplate(w, "base", data)
	case constants.RM_UPLOAD_MODAL:
		template.Must(template.ParseFiles(ptr.UPLOAD_MODAL...)).ExecuteTemplate(w, "test-modal", data)
	case constants.RM_SETTINGS_MODAL:
		template.Must(template.ParseFiles(ptr.SETTINGS_MODAL...)).ExecuteTemplate(w, "settings-modal", data)
	case constants.RM_TABLE:
		template.Must(template.ParseFiles(ptr.TABLE...)).ExecuteTemplate(w, "activity", data)
	case constants.RM_TABLE_REFRESH:
		template.Must(template.ParseFiles(ptr.TABLE...)).ExecuteTemplate(w, "lstable", data)
	case constants.RM_SNIPPET1:
		template.Must(template.ParseFiles(ptr.SNIPPET1...)).ExecuteTemplate(w, "side-nav-categories", data)
	case constants.RM_SNIPPET3:
		template.Must(template.ParseFiles(ptr.SNIPPET3...)).ExecuteTemplate(w, "snippets3", data)
	case constants.RM_PARTIAL1:
		template.Must(template.ParseFiles(ptr.PARTIAL1...)).ExecuteTemplate(w, "partial1", data)
	case constants.RM_CARDS:
		template.Must(template.ParseFiles(ptr.CARDS...)).ExecuteTemplate(w, "cardsvw", data)
	default:
		template.Must(template.ParseFiles(ptr.HOME...)).ExecuteTemplate(w, "base", data)
	}
}
