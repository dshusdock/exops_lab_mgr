package renderview

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/views/base"
	"dshusdock/tw_prac1/internal/views/cardsvw"
	"dshusdock/tw_prac1/internal/views/headervw"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"dshusdock/tw_prac1/internal/views/layoutvw"
	"dshusdock/tw_prac1/internal/views/login"
	"dshusdock/tw_prac1/internal/views/settingsvw"
	"dshusdock/tw_prac1/internal/views/sidenav"
	"fmt"
	"net/http"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]constants.ViewHandler
}

var RenderViewSvc *RenderView

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

type DisplayData struct {
	Base 		base.BaseTemplateparams
	Tmplt   	map[string]*any
	TestStr 	string
}

func InitRouteHandlers() {
	// Register the views
	RenderViewSvc.ViewHandlers["basevw"] = base.AppBaseVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["loginvw"] = login.AppLoginVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["lyoutvw"] = layoutvw.AppLayoutVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["headervw"] = headervw.AppHeaderVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["sidenav"] = sidenav.AppSideNavVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["lstablevw"] = labsystemvw.AppLSTableVW.RegisterHandler()
	RenderViewSvc.ViewHandlers["cardsvw"] = cardsvw.AppCardsVW.RegisterHandler()
	RenderViewSvc.ViewHandlers["settingsvw"] = settingsvw.AppSettingsVw.RegisterHandler()
}

func NewRenderViewSvc(app *config.AppConfig) *RenderView {
	
	obj := &RenderView{
		App: app,
		ViewHandlers: make(map[string]constants.ViewHandler),
	}
	RenderViewSvc = obj

	messagebus.GetBus().Subscribe("Event:Click", RenderViewSvc.HandleMBusRequest)
	return RenderViewSvc
	
}

func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	var rslt any
	var _view int
	
	obj := DisplayData{
		Base: base.BaseTemplateparams{},
		Tmplt: make(map[string]*any),
	}

	if (false) { // some special condition) 
		// do something special that returns rslt
	} else {
		rslt = rv.ViewHandlers[view].HandleRequest(w, r)
	}
	obj.Tmplt[view] = &rslt

	switch view {
	case "loginvw":
		_view = rslt.(login.LoginVwData).View
		obj.Base = rslt.(login.LoginVwData).Base
	case "basevw":		
		_view = rslt.(base.BaseVwData).View
		obj.Base = rslt.(base.BaseVwData).Base
	case "layoutvw":
		_view = rslt.(layoutvw.LayoutVwData).View
		obj.Base = rslt.(layoutvw.LayoutVwData).Base
	case "headervw":
		_view = rslt.(headervw.HeaderVwData).View
		obj.Base = rslt.(headervw.HeaderVwData).Base
	case "sidenav":
		_view = rslt.(sidenav.SideNavVwData).View
		obj.Base = rslt.(sidenav.SideNavVwData).Base
	case "lstablevw":
		_view = rslt.(labsystemvw.LSTableVWData).View
		obj.Base = rslt.(labsystemvw.LSTableVWData).Base
	case "cardsvw":
		_view = rslt.(cardsvw.CardsVwData).View
		obj.Base = rslt.(cardsvw.CardsVwData).Base
		obj.Base.Cards = true
	case "settingsvw":
		_view = rslt.(settingsvw.SettingsVwData).View
		obj.Base = rslt.(settingsvw.SettingsVwData).Base
	default:
	}	

	if _view == constants.RM_NONE { 
		fmt.Println("No view to render")
		return
	}

	// Now go build the view and render it
	rv.RenderTemplate(w, obj, _view)
}

func (rv *RenderView) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[renderview] - HandleMBusRequest")
    d := r.PostForm
	id := d.Get("view_id")	

	switch id {
	case  "headervw":
		rv.HandleHeaderVwEvent(w, r)
		return
	case "sidenav":		
		rv.HandleSideNavEvent(w, r)
	case "settingsvw":
		rv.HandleSettingsVwEvent(w, r)
	}
}

func (rv *RenderView) RenderTemplate(w http.ResponseWriter, data any, view int) {
	render.RenderTemplate_new(w, nil, data, view)
}

func (rv *RenderView) HandleSettingsVwEvent(w http.ResponseWriter, r *http.Request) {
	rv.ProcessRequest(w, r, "settingsvw")
}

func (rv *RenderView) HandleSideNavEvent(w http.ResponseWriter, r *http.Request) {
	rv.ProcessRequest(w, r, "lstablevw")
}

func (rv *RenderView) HandleHeaderVwEvent(w http.ResponseWriter, r *http.Request) {
	d := r.PostForm
	lbl := d.Get("label")

	obj := DisplayData{
		Base: base.GetBaseTemplateObj(),
		Tmplt: make(map[string]*any),
	}
	
	switch lbl{
	case "Table":
		obj.Base.MainTable = true
		obj.Base.Cards = false
		obj.Base.SideNav = true
		
		rslt := rv.ViewHandlers["sidenav"].HandleRequest(w, r)
		obj.Tmplt["sidenav"] = &rslt 
		rslt2 := rv.ViewHandlers["lstablevw"].HandleRequest(w, r)
		obj.Tmplt["lstablevw"] = &rslt2
		rv.RenderTemplate(w, obj, constants.RM_TABLE)
	case "settings": 
		obj.Base.MainTable = false
		obj.Base.Cards = false
		obj.Base.SideNav = false

		rslt := rv.ViewHandlers["settingsvw"].HandleRequest(w, r)
		_view := constants.RM_SETTINGS_MODAL
		obj.Tmplt["settingsvw"] = &rslt
		rv.RenderTemplate(w, obj, _view)
	case "upload":
		obj.Base.MainTable = false
		obj.Base.Cards = false
		obj.Base.SideNav = false

		rslt := rv.ViewHandlers["settingsvw"].HandleRequest(w, r)
		_view := constants.RM_UPLOAD_MODAL
		obj.Tmplt["settingsvw"] = &rslt
		rv.RenderTemplate(w, obj, _view)
	default:
		obj.Base.MainTable = false
		obj.Base.Cards = true
		obj.Base.SideNav = false

		rslt := rv.ViewHandlers["cardsvw"].HandleRequest(w, r)
		_view := constants.RM_TABLE
		obj.Base = rslt.(cardsvw.CardsVwData).Base
		obj.Base.Cards = true
		obj.Tmplt["cardsvw"] = &rslt
		rv.RenderTemplate(w, obj, _view)
	}

}
