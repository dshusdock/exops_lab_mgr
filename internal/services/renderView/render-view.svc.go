package renderview

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"net/http"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewCache 	map[string]*constants.ViewInteface
}

var RenderViewSvc *RenderView

func NewRenderViewSvc(app *config.AppConfig) *RenderView {
	return &RenderView{
		App: app,
		ViewCache: make(map[string]*constants.ViewInteface),
	}
}

func init() {
	RenderViewSvc = &RenderView{
		App: nil,
		ViewCache: make(map[string]*constants.ViewInteface),
	}
}

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

func (rv *RenderView) UpdateView(view constants.ViewInteface) *RenderView{
	return &RenderView{}
}

func (rv *RenderView) RenderTemplate(w http.ResponseWriter, r *http.Request, data any, view int) {
	//rv.App.TemplateCache["loginvw"].Execute(w, nil)
	render.RenderTemplate_new(w, nil, data, view)
}

func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, data any) {

	// c := make(chan BaseVwData)
	// go rv.ViewCache["loginvw"].HandleHttpRequest(w, r, c)
	// rslt := <-c
}

func (rv *RenderView) RegisterView(name string, ptr *constants.ViewInteface) {
	rv.ViewCache[name] = ptr
}