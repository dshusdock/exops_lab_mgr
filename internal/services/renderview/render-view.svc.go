package renderview

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/views/base"
	"dshusdock/tw_prac1/internal/views/headervw"
	"dshusdock/tw_prac1/internal/views/labsystemvw"
	"dshusdock/tw_prac1/internal/views/layoutvw"
	"dshusdock/tw_prac1/internal/views/login"
	"dshusdock/tw_prac1/internal/views/sidenav"
	"fmt"
	"net/http"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]constants.ViewHandler
}

var RenderViewSvc *RenderView

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


	// Repo.App.ViewCache["settingsvw"] = settingsvw.AppSettingsVw.RegisterView(Repo.App) 
	// Repo.App.ViewCache["lstablevw"] = labsystemvw.AppLSTableVW.RegisterView(Repo.App)
	
	// Repo.App.ViewCache["cardsvw"] = cardsvw.AppCardsVW.RegisterView(Repo.App)
	

	// Register the services
	// Repo.App.ViewCache["unigystatus"] = unigystatus.AppStatusSvc.RegisterService(Repo.App)
	// Repo.App.ViewCache["unigydata"] = unigydata.AppUnigyDataSvc.RegisterService(Repo.App)
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

func (rv *RenderView) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[renderview] - HandleMBusRequest")
    d := r.PostForm
	id := d.Get("view_id")	

	switch id {
	case  "headervw":
		rv.HandleHeaderVwRequest(w, r)
	}
}

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
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
		_view = rslt.(*login.LoginVwData).View
		obj.Base = rslt.(*login.LoginVwData).Base
	case "basevw":		
		_view = rslt.(*base.BaseVwData).View
		obj.Base = rslt.(*base.BaseVwData).Base
	case "layoutvw":
		_view = rslt.(*layoutvw.LayoutVwData).View
		obj.Base = rslt.(*layoutvw.LayoutVwData).Base
	case "headervw":
		_view = rslt.(*headervw.HeaderVwData).View
		obj.Base = rslt.(*headervw.HeaderVwData).Base
	case "sidenav":
		_view = rslt.(sidenav.SideNavVwData).View
		obj.Base = rslt.(sidenav.SideNavVwData).Base
	default:
	}	

	if _view == constants.RM_NONE { 
		fmt.Println("No view to render")
		return
	}

	// Now go build the view and render it
	rv.RenderTemplate(w, obj, _view)
}

func (rv *RenderView) RenderTemplate(w http.ResponseWriter, data any, view int) {
	render.RenderTemplate_new(w, nil, data, view)
}

func (rv *RenderView) RegisterView(name string, ptr constants.ViewHandler) {
	rv.ViewHandlers[name] = ptr
}

func (rv *RenderView) HandleHeaderVwRequest(w http.ResponseWriter, r *http.Request) {
	_base := base.GetBaseTemplateObj()
	_base.SideNav = true
	_base.MainTable = true
	_base.LoggedIn = true
	
	obj := DisplayData{
		Base: _base,
		Tmplt: make(map[string]*any),
	}

	obj.TestStr = "Test String"

	rslt := rv.ViewHandlers["sidenav"].HandleRequest(w, r)
	obj.Tmplt["sidenav"] = &rslt 
	// fmt.Printf("%+v\n", rslt)

	rslt2 := rv.ViewHandlers["lstablevw"].HandleRequest(w, r)
	obj.Tmplt["lstablevw"] = &rslt2
	// fmt.Printf("%+v\n", rslt2)

	rv.RenderTemplate(w, obj, constants.RM_TABLE)
	// render.RenderTemplate_new(w, nil, obj, constants.RM_TABLE)
}


////////////////////////////////////////////////



// func (rv *RenderView) UpdateView(view constants.ViewInteface) *RenderView{
// 	return &RenderView{}
// }