package headervw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/views/base"
	"log/slog"
	"strings"

	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/cardsvw"
	// "dshusdock/tw_prac1/internal/views/labsystemvw"

	// "dshusdock/tw_prac1/internal/views/labsystemvw"
	// "dshusdock/tw_prac1/internal/views/tablevw"
	"fmt"
	"net/http"
	"net/url"
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

func (m *HeaderVw) RegisterHandler() constants.ViewHandler {
	return &HeaderVw{}
}

func (m *HeaderVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	slog.Info("Processing request", "ID", "HeaderVw")

	CreateHeaderVwData().ProcessHttpRequest(w, r)
}

func (m *HeaderVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {}

func (m *HeaderVw) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[HeaderVw] - HandleRequest")
	obj := CreateHeaderVwData().ProcessHttpRequest(w, r)

	return obj
}

///////////////////// Header View Data //////////////////////

type HeaderVwData struct {
	Base base.BaseTemplateparams
	Data any
	View int
}

func CreateHeaderVwData() *HeaderVwData {
	return &HeaderVwData{
		Base: base.GetBaseTemplateObj(),
		Data: nil,
		View: constants.RM_HOME,
	}
}

func (m *HeaderVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *HeaderVwData{

	fmt.Printf("[%s] - Processing Http request\n", "HeaderVwData")
	d := r.PostForm
	s := d.Get("label")
	fmt.Println("[HeaderVwData] Recieved Label: ", s)	

	m.ProcessClickEvent(w, r)
	// m.View = constants.RM_HOME
	return m
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
		// This is a composite view
		messagebus.GetBus().Publish("Event:Click", w, r)
		m.View = constants.RM_NONE
		// data := r.PostForm
		// data.Add("event", constants.EVENT_CLICK)
		
		// http.Redirect(w, r, "/element/event/click", http.StatusSeeOther)

		
	case "Cards":
		// messagebus.GetBus().Publish("Event:ViewChange", w, r)
		client := &http.Client{}
    
    	req, _ := http.NewRequest(http.MethodPost, "127.0.0.1/test3", strings.NewReader(url.Values{
        	"key": {"value"},
        	"key1":{"value1"},
    	}.Encode()))

    	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    	client.Do(req)
	}	
}

func (m *HeaderVwData) ToggleView() {
	// if m.ViewFlags[0] {
	// 	m.ViewFlags[0] = false
	// } else {
	// 	m.ViewFlags[0] = true
	// }
}