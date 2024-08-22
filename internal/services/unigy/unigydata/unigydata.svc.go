package unigydata

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/services/database"
	d "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/database/local"
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

func RecordValidDbEndpoints() {
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
			d.CloseUnigyDB()										
		}		
	}	
	slog.Info("LoadUnigyTargets2Db() - Done")		
}

func PopulateZoneInfo() {

	entList := local.GetEnterpriseList()

	for _, ent := range entList {
		// Array of all the nodes in the enterprise
		fmt.Println("Enterprise: ", ent)
		zoneInfoAry := getZoneInfo(ent.Data[0])
		if len(zoneInfoAry) == 0 {
			continue
		}
		for _, zoneInfo := range zoneInfoAry {
			fmt.Println("ZID:" + zoneInfo.Zid)
			fmt.Println("VIP:" + zoneInfo.Vip)
			fmt.Println("CCM1:" + zoneInfo.Ccm1)
			fmt.Println("CCM2:" + zoneInfo.Ccm2)


			// Write the zone info to the local database
			local.WriteZoneInfoData(zoneInfo)
		}
	}
}

func getZoneInfo(ent string) []con.ZoneInfo {
	zi := []con.ZoneInfo{}

	target := getDBEndpoint(ent)
	if target == "no endpoint" {
		return []con.ZoneInfo{}
	}

	err := d.ConnectUnigyDB(target)
	if err != nil {}

	s := fmt.Sprintf(q.SQL_QUERIES_UNIGY["QUERY_1"].Qry )
	da := d.ReadUnigyDBwithType[q.TBL_NZData](s)
	
	for _, el := range da {
		// fmt.Println(el.Data[3])

		zi = append(zi, con.ZoneInfo{
			Enterprise: ent,
			Zid: el.Data[3],
			Vip: el.Data[2],
			Ccm1: el.Data[0],
			Ccm2: el.Data[1],	
			Online: true,
			Status: "active",			

		})

	}
	return zi
}

func getDBEndpoint(ent string) string {
	xx := `select targetIP from UnigyDatabaseTargets where enterprise="%s" and status="available" limit 1`
	s := fmt.Sprintf(xx, ent) 
	
	rslt := d.ReadLocalDBwithType[q.StringVal](s)
	if len(rslt) == 0 {
		return "no endpoint"
	}
	return rslt[0].Data[0]
}

//select targetIP from UnigyDatabaseTargets where enterprise="Sleepy" and status="available" limit 1;

//me.txt
//mycluster.txt
//haservices/checkBasicStatus
//ums/about.xml
