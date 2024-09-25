package unigydata

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	db "dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/local"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/unigy/unigystatus"
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

// LoadUnigyTargets2Db - This function will load the Unigy targets into the local database
// It will check if the target is reachable i.e. SQL connection to Unigy DB and mark it as 
// available or unavailable in the UniigyDatabaseTargets table
func IdentifyValidDbEndpoints() {
	//  Get the list of enterprise names
	entList, _ := db.ReadDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)
	for _, ent := range entList {	
		// Get a list of IP's based on the enterprise name	
		str := fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_7"].Qry + "\"%s\"\n", ent.Data[0])
		ipList, _ := db.ReadDBwithType[q.TBL_ServerTypeList](str)

		for _, ip := range ipList {
			slog.Info("LoadUnigyTargets2Db() - Checking IP", "IP", ip.Data[0])
			// Check if the IP is reachable
			err := db.ConnectUnigyDB(ip.Data[0])
			if err != nil {				
				// Mark as inactive
				wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "unavailable")	
				db.WriteLocalDB(wrStr)								
				continue
			} 
			// Mark as active
			wrStr := fmt.Sprintf("INSERT into UnigyDatabaseTargets values (\"%s\", \"%s\", \"%s\")", ent.Data[0], ip.Data[0], "available")
			db.WriteLocalDB(wrStr)
			db.CloseUnigyDB()										
		}		
	}	
	slog.Info("LoadUnigyTargets2Db() - Done")		
}

// PopulateZoneInfoTable - This function will populate the local database "ZoneInfo" table with the zone information
// received from the apppropiate Unigy Enterprise database
func PopulateZoneInfoTable() {

	entList, _ := local.GetEnterpriseList()

	for _, ent := range entList {
		// Array of all the nodes in the enterprise
		zoneInfoAry := getEnterpriseZoneInfo(ent.Data[0])
		if len(zoneInfoAry) == 0 {
			continue
		}
		for _, zoneInfo := range zoneInfoAry {
			// Write the zone info to the local database
			fmt.Println("Writing ZoneInfo: ", zoneInfo)
			local.WriteZoneInfoData(zoneInfo)
		}
	}
}

// Helper functions

// getEnterpriseZoneInfo - This function will get the zone information for the enterprise
func getEnterpriseZoneInfo(ent string) []con.ZoneInfo {
	zi := []con.ZoneInfo{}

	// Get the target IP for the enterprise
	target := getAvailableDBEndpoint(ent)
	if target == "no endpoint" {
		return []con.ZoneInfo{}
	}

	// Connect to the Unigy enterprise database
	err := db.ConnectUnigyDB(target)
	if err != nil {}

	// Read the newZoneData table for the enterprise
	s := fmt.Sprintf(q.SQL_QUERIES_UNIGY["QUERY_1"].Qry )
	da, _ := db.ReadUnigyDBwithType[q.TBL_NZData](s)

	for _, el := range da {
		
		r1, _ := unigystatus.GetServerState(el.Data[0])
		r2, _ := unigystatus.GetServerState(el.Data[1])

		fmt.Println("Server state for ccm1: ", r1, el.Data[0])
		fmt.Println("Server state for ccm2: ", r2, el.Data[1])

		zi = append(zi, con.ZoneInfo{
			Enterprise: ent,
			Zid: el.Data[3],
			Vip: el.Data[2],
			Ccm1: con.Server{
				IP:  el.Data[0],
				SWVersion: "",
				State: r1,
				Active: false,
				Standby: false,				
			},
			Ccm2: con.Server{
				IP:  el.Data[1],
				SWVersion: "",
				State: r2,
				Active: false,
				Standby: false,				
			},	
			Online: true,
			Status: "active",			
		})
	}
	return zi
}

// getAvailableDBEndpoint - This function will get the SQL DB endpoint for the specified Unigy enterprise
func getAvailableDBEndpoint(ent string) string {
	xx := `select targetIP from UnigyDatabaseTargets where enterprise="%s" and status="available" limit 1`
	s := fmt.Sprintf(xx, ent) 
	
	rslt, _ := db.ReadDBwithType[q.StringVal](s)
	if len(rslt) == 0 {
		return "no endpoint"
	}
	return rslt[0].Data[0]
}

func PopulateDeviceTableByEnterprise(ent string) {
	// Get the target IP for the enterprise
	target := getAvailableDBEndpoint(ent)
	if target == "no endpoint" {
		return
	}

	// Connect to the Unigy enterprise database
	err := db.ConnectUnigyDB(target)
	if err != nil {
		return
	}

	// Read the LabSystem table for the enterprise
	da, _ := db.ReadUnigyDBwithType[q.UNIGY_TBL_DEVICE](q.SQL_QUERIES_UNIGY["QUERY_2"].Qry)

	for _, el := range da {

		p := q.UNIGY_TBL_DEVICE{}
		
		switch el.Data[1] {
		case "3":
			el.Data[1] = "max"
		case "4":
			el.Data[1] = "UDA"
		case "6":
			el.Data[1] = "pulse"
		case "7":
			el.Data[1] = "mercury"
		}
		p.DeviceState = el.Data[0]
		p.DeviceTypeId = el.Data[1]
		p.Equipped = el.Data[2]
		p.DunkinLocationId = el.Data[3]
		p.IPAddress = el.Data[4]
		p.MacAddress = el.Data[5]
		p.ParentZoneId = el.Data[6]

		fmt.Println("After Writing MdcData: ", el, ent)
		local.WriteDeviceData(p, ent)
	}
}	