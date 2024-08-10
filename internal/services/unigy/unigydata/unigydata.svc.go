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

func LoadUnigyTargets2Db() {
	
	//  Get the list of enterprise names
	entList := d.ReadLocalDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
	for _, ent := range entList {	
		// Get a list of IP's based on the enterprise name	
		str := fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_7"].Qry + "\"%s\"\n", ent.Data[0])
		ipList := d.ReadLocalDBwithType[q.TBL_ServerTypeList](str)

		for _, ip := range ipList {
			slog.Info("LoadUnigyTargets2Db() - Checking IP", "IP", ip.Data[0])
			// Check if the IP is reachable
			err := d.ConnectUnigyDB(ip.Data[0])
			if err != nil {				
				// Mark as inactive
				wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "unavailable")	
				database.WriteLocalDB(wrStr)								
				continue
			} 
			// Mark as active
			wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "available")
			database.WriteLocalDB(wrStr)										
		}		
	}	
	slog.Info("LoadUnigyTargets2Db() - Done")		
}

//me.txt
//mycluster.txt
//haservices/checkBasicStatus
//ums/about.xml
