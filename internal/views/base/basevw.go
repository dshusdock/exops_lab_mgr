package base

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/render"
	"fmt"
	"log"
	"net/http"
)

type BaseTemplateparams struct {
	LoggedIn 					bool
	DisplayLogin  				bool
	DisplayCreateAccount 		bool
	DisplayCreatAcctResponse 	bool
	SideNav	      				bool
	MainTable	  				bool
	Cards		  				bool
}

func GetBaseTemplateObj() BaseTemplateparams{
	return BaseTemplateparams{
		LoggedIn: false,
		DisplayLogin: true,
		DisplayCreateAccount: false,
		DisplayCreatAcctResponse: false,
		SideNav: false,
		MainTable: false,
		Cards: false,
	}
}

///////////////////// Base View //////////////////////
type BaseVw struct {
	App *config.AppConfig
}

var AppBaseVw *BaseVw

func init() {
	AppBaseVw = &BaseVw{
		App: nil,
	}
	// renderview.RenderViewSvc.RegisterView("basevw", AppBaseVw)
}

func (m *BaseVw) RegisterView(app *config.AppConfig) *BaseVw{
	log.Println("Registering AppLayoutVw...")
	AppBaseVw.App = app
	return AppBaseVw
}

func (m *BaseVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[lyoutvw] - Processing request")
	CreateBaseVwData().ProcessHttpRequest(w, r)
}

func (m *BaseVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	CreateBaseVwData().ProcessMBusRequest(w, r)
}

func (m *BaseVw) HandleRequest(w http.ResponseWriter, r *http.Request, c chan BaseVwData) {
	fmt.Println("[lyoutvw] - HandleRequest")
	// rslt := CreateBaseVwData().ProcessHttpRequest(w, r)
	// c <- rslt
}


///////////////////// Base View Data //////////////////////

type BaseVwData struct {
	Base BaseTemplateparams
	Data any
}

func CreateBaseVwData() *BaseVwData {
	return &BaseVwData{
		Base: GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *BaseVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[lyoutvw] - Processing request")
	render.RenderModal(w, nil, m)
	
}

func (m *BaseVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {

}