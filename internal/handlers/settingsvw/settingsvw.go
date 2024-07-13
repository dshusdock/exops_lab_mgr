package settingsvw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/datetime"
	"fmt"
	"net/http"
	"net/url"
)

type AppSettingsVwData struct {
	Lbl string
}

type SettingsVw struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppSettingsVw *SettingsVw

func init() {
	AppSettingsVw = &SettingsVw{
		Id:         "settingsvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *SettingsVw) RegisterView(app config.AppConfig) *SettingsVw{
	AppSettingsVw.App = &app
	return AppSettingsVw
}

func (m *SettingsVw) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[settingsvw] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "Test Button":
		fmt.Println("Test Button Clicked")

		datetime.Prac1()
		datetime.Prac2()
		datetime.Prac3()
	}
	
	
}

