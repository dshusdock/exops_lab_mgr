package config

import (
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/services/database"
	"log"
	"text/template"

	"github.com/asaskevich/EventBus"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	ViewCache     map[string]constants.ViewInteface
	DBA 		  *database.DBAccess
	Bus 		  EventBus.Bus
	SideNav	      bool
	MainTable	  bool
}
