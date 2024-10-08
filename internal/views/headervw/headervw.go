package headervw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/views/base"
	"log/slog"

	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/labsystemvw"

	// "dshusdock/tw_prac1/internal/views/labsystemvw"
	// "dshusdock/tw_prac1/internal/views/tablevw"
	"fmt"
	"net/http"
)

type HeaderVw struct {
	App  *config.AppConfig
}

var AppHeaderVw *HeaderVw

type AppHdrVwData struct {
	Lbl string
}

func init() {
	AppHeaderVw = &HeaderVw{
		App: nil,
	}
	messagebus.GetBus().Subscribe("Event:ViewChange", AppHeaderVw.HandleMBusRequest)
}

func (m *HeaderVw) RegisterView(app *config.AppConfig) *HeaderVw {
	slog.Info("Registering AppHeaderVw...")
	AppHeaderVw.App = app
	return AppHeaderVw
}

func (m *HeaderVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	slog.Info("Processing request", "ID", "HeaderVw")

	CreateHeaderVwData().ProcessHttpRequest(w, r)
}

func (m *HeaderVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {}

///////////////////// Header View Data //////////////////////

type HeaderVwData struct {
	Base base.BaseTemplateparams
	Data any
}

func CreateHeaderVwData() *HeaderVwData {
	return &HeaderVwData{
		Base: base.GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *HeaderVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[%s] - Processing Http request\n", "HeaderVwData")
	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)	

	m.ProcessClickEvent(w, r)
}

func (m *HeaderVwData) ProcessMbusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] - Processing MBus request\n", "HeaderVwData")
	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)	
}

func (m *HeaderVwData) ProcessClickEvent(w http.ResponseWriter, r *http.Request) {
	d := r.PostForm
	
	lbl := d.Get("label")
	slog.Info("ProcessClickEvent - ", "ID", lbl)

	switch lbl {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "settings":
		render.RenderTemplate_new(w, nil, m.Base, constants.RM_SETTINGS_MODAL)
	case "Table":
		messagebus.GetBus().Publish("Event:Click", w, r)
	case "Cards":
		messagebus.GetBus().Publish("Event:ViewChange", w, r)
	}	
}

func (m *HeaderVwData) ToggleView() {
	// if m.ViewFlags[0] {
	// 	m.ViewFlags[0] = false
	// } else {
	// 	m.ViewFlags[0] = true
	// }
}