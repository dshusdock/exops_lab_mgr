package renderview

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/views/base"
	"dshusdock/tw_prac1/internal/views/login"
	"fmt"
	"net/http"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]constants.ViewHandler
}

var RenderViewSvc *RenderView

func InitRouteHandlers() {
	// Register the views
	RenderViewSvc.ViewHandlers["basevw"] = base.AppBaseVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["loginvw"] = login.AppLoginVw.RegisterHandler()

	
	// Repo.App.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(Repo.App)
	
	// Repo.App.ViewCache["headervw"] = headervw.AppHeaderVw.RegisterView(Repo.App)
	// Repo.App.ViewCache["sidenav"] = sidenav.AppSideNavVw.RegisterView(Repo.App)	

	// Repo.App.ViewCache["settingsvw"] = settingsvw.AppSettingsVw.RegisterView(Repo.App) 
	// Repo.App.ViewCache["lstablevw"] = labsystemvw.AppLSTableVW.RegisterView(Repo.App)
	
	// Repo.App.ViewCache["cardsvw"] = cardsvw.AppCardsVW.RegisterView(Repo.App)
	

	// Register the services
	// Repo.App.ViewCache["unigystatus"] = unigystatus.AppStatusSvc.RegisterService(Repo.App)
	// Repo.App.ViewCache["unigydata"] = unigydata.AppUnigyDataSvc.RegisterService(Repo.App)
}

func NewRenderViewSvc(app *config.AppConfig) *RenderView {
	return &RenderView{
		App: app,
		ViewHandlers: make(map[string]constants.ViewHandler),
	}
}

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	
	// c := make(chan any)
	// d := make(chan int)
	var rslt any


	switch view {
	case "loginvw":
		fmt.Println("Processing loginvw")
		rslt = rv.ViewHandlers[view].HandleRequest(w, r)
		
	// case "basevw":
	// 	go rv.ViewHandlers[view].HandleRequest(w, r, c, d)
	// case "layoutvw":
	// 	go rv.ViewHandlers[view].HandleRequest(w, r, c, d)
	// case "lstablevw":
	// 	go rv.ViewHandlers[view].HandleRequest(w, r, c, d)

	default:
	}
	// rslt := <- c
	// _view := <- d

	// fmt.Println("ProcessRequest: ", rslt)
	fmt.Printf(" ProcessRequest: %+v\n", rslt)
	// Now go build the view and render it
	// rv.RenderTemplate(w, r, rslt, view)
}

func (rv *RenderView) RenderTemplate(w http.ResponseWriter, r *http.Request, data any, view int) {
	//rv.App.TemplateCache["loginvw"].Execute(w, nil)
	render.RenderTemplate_new(w, nil, data, view)
}

func (rv *RenderView) RegisterView(name string, ptr constants.ViewHandler) {
	rv.ViewHandlers[name] = ptr
}


////////////////////////////////////////////////



// func (rv *RenderView) UpdateView(view constants.ViewInteface) *RenderView{
// 	return &RenderView{}
// }