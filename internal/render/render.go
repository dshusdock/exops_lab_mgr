package render

import (
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
