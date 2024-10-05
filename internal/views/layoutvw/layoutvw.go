package layoutvw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/render"
	// "dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
	// "net/url"
)

type AppLytVwData struct {
	Lbl string
}

type LayoutVw struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppLayoutVw *LayoutVw

func init() {
	AppLayoutVw = &LayoutVw{
		Id:         "lyoutvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *LayoutVw) RegisterView(app config.AppConfig) *LayoutVw{
	log.Println("Registering AppLayoutVw...")
	// messagebus.GetBus().Subscribe("Event:Change", AppLayoutVw.ProcessChangEvent)
	AppLayoutVw.App = &app
	return AppLayoutVw
}

func (m *LayoutVw) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[lyoutvw] - Processing request")
	render.RenderModal(w, nil, nil)
}

func (m *LayoutVw) ProcessChangEvent() {
	fmt.Println("\n[lyoutvw] Process Change Event")

}

func (m *LayoutVw) ToggleView() {
	if m.ViewFlags[0] {
		m.ViewFlags[0] = false
	} else {
		m.ViewFlags[0] = true
	}
}