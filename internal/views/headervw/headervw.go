package headervw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"log/slog"

	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/labsystemvw"

	// "dshusdock/tw_prac1/internal/views/labsystemvw"
	// "dshusdock/tw_prac1/internal/views/tablevw"
	"fmt"
	"net/http"
	"net/url"
)

type AppHdrVwData struct {
	Lbl string
}

type HeaderVw struct {
	App        *config.AppConfig
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
		Data:       "",
		Htmx:       nil,
	}
	messagebus.GetBus().Subscribe("Event:ViewChange", AppHeaderVw.ProcessInternalRequest)
}

func (m *HeaderVw) RegisterView(app config.AppConfig) *HeaderVw {
	slog.Info("Registering AppHeaderVw...")
	AppHeaderVw.App = &app
	return AppHeaderVw
}

func (m *HeaderVw) ProcessRequest(w http.ResponseWriter, d url.Values) {

	slog.Info("Processing request", "ID", m.Id)
	s := d.Get("label")
	slog.Info("Incoming: ", "Label", s)

	m.ProcessClickEvent(w, d)
}

func (m *HeaderVw) ProcessInternalRequest(w http.ResponseWriter, d url.Values) {

	fmt.Printf("[%s] - Processing Internal request\n", m.Id)
	s := d.Get("label")
	fmt.Println("Label: ", s)	
}

func (m *HeaderVw) ProcessClickEvent(w http.ResponseWriter, d  url.Values) {
	if d.Get("view_id") != m.Id {return}
	lbl := d.Get("label")
	slog.Info("ProcessClickEvent - ", "ID", lbl)

	switch lbl {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "settings":
		render.RenderTemplate_new(w, nil, m.App, constants.RM_SETTINGS_MODAL)
	case "Table":
		messagebus.GetBus().Publish("Event:Click", w, d)
	case "Cards":
		messagebus.GetBus().Publish("Event:ViewChange", w, d)
	}	
}

func (m *HeaderVw) ToggleView() {
	if m.ViewFlags[0] {
		m.ViewFlags[0] = false
	} else {
		m.ViewFlags[0] = true
	}
}