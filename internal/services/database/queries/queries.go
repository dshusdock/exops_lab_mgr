package queries

import (
	"reflect"
)
var SQL_QUERIES_LOCAL = make (map[string]lsquery)
var SQL_QUERIES_UNIGY = make (map[string]lsquery)


type lsquery struct {
	Qry string
	model interface{}
}

func init() {
	
	// LOCAL DATABASE
	SQL_QUERIES_LOCAL["QUERY_1"] = lsquery{"select unique enterprise from LabSystem", reflect.TypeOf(TBL_EnterpriseList{})}
	SQL_QUERIES_LOCAL["QUERY_2"] = lsquery{"select * from LabSystem where enterprise = ?", reflect.TypeOf(MdcData{})}	
	SQL_QUERIES_LOCAL["QUERY_3"] = lsquery{`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`, reflect.TypeOf(DataVw1{})}
	SQL_QUERIES_LOCAL["QUERY_4"] = lsquery{"select unique swVer from LabSystem", reflect.TypeOf(TBL_SWVerList{})}
	SQL_QUERIES_LOCAL["QUERY_5"] = lsquery{`select unique enterprise from LabSystem where role="Unigy"`, reflect.TypeOf(TBL_EnterpriseList{})}
	SQL_QUERIES_LOCAL["QUERY_6"] = lsquery{`select unique serverType from LabSystem where enterprise = `, reflect.TypeOf(TBL_ServerTypeList{})}
	SQL_QUERIES_LOCAL["QUERY_7"] = lsquery{`select unique iPAddress from LabSystem where enterprise = `, reflect.TypeOf(TBL_CcmIPList{})}
	SQL_QUERIES_UNIGY["QUERY_8"] = lsquery{`select targetIP from UnigyDatabaseTargets where enterprise=%s and status="available" limit 1`, reflect.TypeOf(StringVal{})}
	SQL_QUERIES_LOCAL["QUERY_9"] = lsquery{`select * from ZoneInfo where enterprise= `, reflect.TypeOf(TBL_ServerTypeList{})}
	SQL_QUERIES_LOCAL["QUERY_10"] = lsquery{`select username from User`, reflect.TypeOf(TBL_UserNames{})}
	SQL_QUERIES_LOCAL["QUERY_11"] = lsquery{`select username, password from User where username= `, reflect.TypeOf(TBL_UserInfo{})}



	// UNIGY DATABASE
	SQL_QUERIES_UNIGY["QUERY_1"] = lsquery{`select server1,server2,vip,zid from NewZoneData`, reflect.TypeOf(TBL_NZData{})}
	SQL_QUERIES_UNIGY["QUERY_2"] = lsquery{`SELECT deviceState,
       deviceTypeId,
       equipped,
       dunkinLocationId,
       iPAddress,
       macAddress,
       parentZoneId
  FROM dunkin.Device Device
 WHERE (deviceTypeId = 3 or deviceTypeId = 4 or deviceTypeId = 6 or deviceTypeId = 7)`, reflect.TypeOf(UNIGY_TBL_DEVICE{})}
}

// LOCAL DATABASE

type TBL_UserInfo struct {
	Username string
	Password string
}

type TBL_UserNames struct {
	Username string
}

type TBL_EnterpriseList struct {
	Enterprise string	
}

type TBL_CcmIPList struct {
	iPAddress string	
}

type TBL_ServerTypeList struct {
	ServerType string	
}

type TBL_SWVerList struct {
	SWVer string	
}

type TBL_ZoneInfo struct {
	
}

type UNIGY_TBL_DEVICE struct {
	DeviceState 		string
	DeviceTypeId 		string
	Equipped			string
	DunkinLocationId 	string
	IPAddress 			string
	MacAddress 			string
	ParentZoneId 		string
}

type DataVw1 struct {
	Vip        string
	SwVer      string
	Enterprise string
	Name       string
}

type MdcData struct {
	Cab               string
	CabULocation      string
	Iso               string
	Name              string
	SerialNbr         string
	IPAddress         string
	Vip               string
	IdracIp           string
	SwVer             string
	ServerType        string
	Enterprise        string
	Role              string
	Comments          string
	VmLabServerHostIp string
}

type StringVal struct {
	Val	string
}

// UNIGY DATABASE
type TBL_NZData struct {
	Server1 string
	Server2 string
	Vip     string
	Zid     string
}

// type labsystem_queries struct {
// 	QUERY_1 lsquery
// 	QUERY_2 lsquery
// 	QUERY_3 lsquery
// 	QUERY_4 lsquery
// 	QUERY_5 lsquery
// 	QUERY_6 lsquery
// 	QUERY_7 lsquery
// }

// func TBL_LAB_SYSTEM_QRY() *labsystem_queries {
// 	return &labsystem_queries{
// 		QUERY_1: lsquery{"select unique enterprise from LabSystem", reflect.TypeOf(TBL_EnterpriseList{})},
// 		QUERY_2: lsquery{"select * from LabSystem where enterprise = ?", reflect.TypeOf(MdcData{})},
// 		QUERY_3: lsquery{`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`, reflect.TypeOf(DataVw1{})},	
// 		QUERY_4: lsquery{"select unique swVer from LabSystem", reflect.TypeOf(TBL_SWVerList{})},
// 		QUERY_5: lsquery{`select unique enterprise from LabSystem where role="Unigy"`, reflect.TypeOf(TBL_EnterpriseList{})},
// 		QUERY_6: lsquery{`select unique serverType from LabSystem where enterprise = `, reflect.TypeOf(TBL_ServerTypeList{})},
// 		QUERY_7: lsquery{`select server1,server2,vip,zid from NewZoneData`, reflect.TypeOf(TBL_NZData{})},	
// 	}
// }