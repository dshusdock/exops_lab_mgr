package database

import "reflect"

type lsquery struct {
	Qry string
	model interface{}
}

type labsystem_queries struct {
	QUERY_1 lsquery
	QUERY_2 lsquery
	QUERY_3 lsquery
	QUERY_4 lsquery
	QUERY_5 lsquery
	QUERY_6 lsquery
	QUERY_7 lsquery
}

func TBL_LAB_SYSTEM_QRY() *labsystem_queries {
	return &labsystem_queries{
		QUERY_1: lsquery{"select unique enterprise from LabSystem", reflect.TypeOf(TBL_EnterpriseList{})},
		QUERY_2: lsquery{"select * from LabSystem where enterprise = ?", reflect.TypeOf(MdcData{})},
		QUERY_3: lsquery{`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`, reflect.TypeOf(DataVw1{})},	
		QUERY_4: lsquery{"select unique swVer from LabSystem", reflect.TypeOf(TBL_SWVerList{})},
		QUERY_5: lsquery{`select unique enterprise from LabSystem where role="Unigy"`, reflect.TypeOf(TBL_EnterpriseList{})},
		QUERY_6: lsquery{`select unique serverType from LabSystem where enterprise = `, reflect.TypeOf(TBL_ServerTypeList{})},
		QUERY_7: lsquery{`select server1,server2,vip,zid from NewZoneData`, reflect.TypeOf(TBL_NZData{})},	
	}
}

type TBL_NZData struct {
	Server1 string
	Server2 string
	Vip     string
	Zid     string
}

type TBL_EnterpriseList struct {
	Enterprise string	
}

type TBL_ServerTypeList struct {
	ServerType string	
}

type TBL_SWVerList struct {
	SWVer string	
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
