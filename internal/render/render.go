package render

import (
	"dshusdock/tw_prac1/internal/constants"
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

	switch _type {
	case constants.RM_HOME:
		template.Must(template.ParseFiles(ptr.HOME...)).ExecuteTemplate(w, "base", data)
	case constants.RM_ADD_FORM:
		template.Must(template.ParseFiles(ptr.ADD_FORM...)).ExecuteTemplate(w, "base", data)
	default:
	}
}
