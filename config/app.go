package config

import (
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/services/database"
	"log"
	"text/template"
	"github.com/alexedwards/scs/v2"
	"github.com/asaskevich/EventBus"
)

type AppConfig struct {
	UseCache      			bool
	TemplateCache 			map[string]*template.Template
	InfoLog       			*log.Logger
	InProduction  			bool
	ViewCache     			map[string]constants.ViewInteface
	DBA 		  			*database.LocalDBAccess
	Bus 		  			EventBus.Bus
	SideNav	      			bool
	MainTable	  			bool
	Cards		  			bool
	SessionManager 	 		*scs.SessionManager
}


