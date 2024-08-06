package database

import (
	"fmt"
	"reflect"
)

func GetTableData() ([]DataVw1, error) {
	// An albums slice to hold data from returned rows.
	var data []DataVw1

	rows := DBA.Read(TBL_LAB_SYSTEM_QRY().QUERY_3.Qry)

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dv DataVw1
		if err := rows.Scan(&dv.Vip, &dv.Enterprise, &dv.Name, &dv.SwVer); err != nil {
			return nil, fmt.Errorf("ERROR: %v", err)
		}
		data = append(data, dv)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}
	return data, nil
}

func PrintTableData() {
	fmt.Println("Printing Table Data")
	data, err := GetTableData()
	fmt.Println("Data: ", len(data))
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}
	for _, dv := range data {
		fmt.Printf("%s %s %s %s\n", dv.Vip, dv.SwVer, dv.Enterprise, dv.Name)
	}
}

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
}

func TBL_LAB_SYSTEM_QRY() *labsystem_queries {
	return &labsystem_queries{
		QUERY_1: lsquery{"select unique enterprise from LabSystem", reflect.TypeOf(TBL_EnterpriseList{})},
		QUERY_2: lsquery{"select * from LabSystem where enterprise = ?", reflect.TypeOf(MdcData{})},
		QUERY_3: lsquery{`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`, reflect.TypeOf(DataVw1{})},	
		QUERY_4: lsquery{"select unique swVer from LabSystem", reflect.TypeOf(TBL_SWVerList{})},
		QUERY_5: lsquery{`select unique enterprise from LabSystem where role="Unigy"`, reflect.TypeOf(TBL_EnterpriseList{})},
		QUERY_6: lsquery{`select unique serverType from LabSystem where enterprise = `, reflect.TypeOf(TBL_ServerTypeList{})},	
	}
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
