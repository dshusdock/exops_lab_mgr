package LSTableVW

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"fmt"
	"net/http"
	"net/url"
)

type AppLSTableVWData struct {
	Lbl string
}

type LSTableVW struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppLSTableVW *LSTableVW

func init() {
	AppLSTableVW = &LSTableVW{
		Id:         "LSTableVW",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *LSTableVW) RegisterView(app config.AppConfig) *LSTableVW{
	AppLSTableVW.App = &app
	return AppLSTableVW
}

func (m *LSTableVW) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[LSTableVW] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	}
	
	
}

