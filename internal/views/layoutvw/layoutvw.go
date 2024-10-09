package layoutvw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/renderview"
	b "dshusdock/tw_prac1/internal/views/base"

	// "dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
	// "net/url"
)

type LayoutVw struct {
	App *config.AppConfig
}

var AppLayoutVw *LayoutVw

func init() {
	AppLayoutVw = &LayoutVw{
		App: nil,
	}
	messagebus.GetBus().Subscribe("Event:ViewChange", AppLayoutVw.HandleMBusRequest)
}

func (m *LayoutVw) RegisterView(app *config.AppConfig) *LayoutVw{
	log.Println("Registering AppLayoutVw...")
	AppLayoutVw.App = app
	return AppLayoutVw
}

func (m *LayoutVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[lyoutvw] - Processing request")
	CreateLayoutVwData().ProcessHttpRequest(w, r)

	// render.RenderModal(w, nil, nil)
}

func (m *LayoutVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	CreateLayoutVwData().ProcessMBusRequest(w, r)
}

///////////////////// Layout View Data //////////////////////

type LayoutVwData struct {
	Base b.BaseTemplateparams
	Data any
}

type AppLytVwData struct {
	Lbl string
}

func CreateLayoutVwData() *LayoutVwData {
	return &LayoutVwData{
		Base: b.GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *LayoutVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) {

	renderview.RenderViewSvc.RenderTemplate(w, r, m.Data, 0)
	// render.RenderTemplate_new(w, nil, m.Data, 0)
}

func (m *LayoutVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {}
