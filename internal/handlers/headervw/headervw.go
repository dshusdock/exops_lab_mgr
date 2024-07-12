package headervw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"fmt"
	"net/http"
	"net/url"
)

type AppHdrVwData struct {
	Lbl string
}

type HeaderVw struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppHeaderVw *HeaderVw

func init() {
	AppHeaderVw = &HeaderVw{
		Id:         "headervw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *HeaderVw) RegisterView(app config.AppConfig) *HeaderVw{
	AppHeaderVw.App = &app
	return AppHeaderVw
}

func (m *HeaderVw) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[headervw] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "settings":
		render.RenderTemplate_new(w, nil, nil, constants.RM_SETTINGS_MODAL)

	}
	
}

func (m *HeaderVw) ToggleView() {
	if m.ViewFlags[0] {
		m.ViewFlags[0] = false
	} else {
		m.ViewFlags[0] = true
	}
}