package unigydata

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/services/database"
	d "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

type UnigyDataSvc struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppUnigyDataSvc *UnigyDataSvc

func init() {
	AppUnigyDataSvc = &UnigyDataSvc{
		Id:         "unigydatasvc",
		Data: "",
	}
	messagebus.GetBus().Subscribe("UnigyDataSvc:Request", AppUnigyDataSvc.ProcessMBusRequest)
}

func (m *UnigyDataSvc) RegisterService(app config.AppConfig) *UnigyDataSvc{
	slog.Info("Registering AppStatusSvc...")
	AppUnigyDataSvc.App = &app
	return AppUnigyDataSvc
}

func (m *UnigyDataSvc) ProcessMBusRequest() {
	slog.Info("Processing MBus Request")
}

func (m *UnigyDataSvc) ProcessRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("Processing request", "ID", m.Id)
}

func LoadZoneData() {

	entList := d.ReadLocalDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
	for _, ent := range entList {
		//  1 - Get a list of IP's based on the enterprise name
		str := fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_7"].Qry + "\"%s\"\n", ent.Data[0])
		ipList := d.ReadLocalDBwithType[q.TBL_ServerTypeList](str)
		count := 0
		for _, ip := range ipList {
			// Check if the IP is reachable
			err := d.ConnectUnigyDB(ip.Data[0])
			if err != nil {
				count++
				if count > 3 {
					// Mark as inactive
					wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "inactive")	
					database.WriteLocalDB(wrStr)				
					break
				}
				continue
			} 
			// Mark as active
			wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "active")
			database.WriteLocalDB(wrStr)
				
			
			
		}		
	}
		
	
}