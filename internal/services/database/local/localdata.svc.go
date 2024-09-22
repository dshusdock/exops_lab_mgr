package local

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

type LocalDataSvc struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppLocalDataSvc *LocalDataSvc

func init() {
	AppLocalDataSvc = &LocalDataSvc{
		Id:         "localdatasvc",
		Data: "",
	}
	messagebus.GetBus().Subscribe("LocalDataSvc:Request", AppLocalDataSvc.ProcessMBusRequest)
}

func (m *LocalDataSvc) RegisterService(app config.AppConfig) *LocalDataSvc{
	slog.Info("Registering AppStatusSvc...")
	AppLocalDataSvc.App = &app
	return AppLocalDataSvc
}

func (m *LocalDataSvc) ProcessMBusRequest() {
	slog.Info("Processing MBus Request")
}

func (m *LocalDataSvc) ProcessRequest(w http.ResponseWriter, d url.Values) {
	slog.Info("Processing request", "ID", m.Id)
}




1func GetEnterpriseList() ([]con.RowData, error){
	return d.ReadDBwithType[q.TBL_EnterpriseList](q.SQL_QUERIES_LOCAL["QUERY_5"].Qry)	
}

func GetUserNames() ([]con.RowData, error){
	return  d.ReadDBwithType[q.TBL_UserNames](q.SQL_QUERIES_LOCAL["QUERY_10"].Qry)	
}

func GetUserInfo(name string) ([]con.RowData, error){
	
	s :=fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_11"].Qry + "\"%s\"", name)
	fmt.Println("SQL: ", s)
	return d.ReadDBwithType[q.TBL_UserInfo](s)
}

func WriteZoneInfoData(z con.ZoneInfo) {
			
	str := fmt.Sprintf(`INSERT into ZoneInfo (enterprise, zid, vip, ccm1, ccm2, ccm1_state, ccm2_state, online, status) values("%s","%s","%s","%s","%s","%s","%s","%v","%s")`, 
		z.Enterprise, z.Zid, z.Vip, z.Ccm1.IP, z.Ccm2.IP, z.Ccm1.State, z.Ccm2.State, z.Online, z.Status)

		d.WriteLocalDB(str)
}



