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
		fmt.Print("Reading Data")
		var dv DataVw1
		if err := rows.Scan(&dv.vip, &dv.enterprise, &dv.name, &dv.swVer); err != nil {
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
		fmt.Printf("%s %s %s %s\n", dv.vip, dv.swVer, dv.enterprise, dv.name)
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
}

func TBL_LAB_SYSTEM_QRY() *labsystem_queries {
	return &labsystem_queries{
		QUERY_1: lsquery{"select unique enterprise from LabSystem", reflect.TypeOf(TBL_EnterpriseList{})},
		QUERY_2: lsquery{"select * from LabSystem where enterprise = ?", reflect.TypeOf(MdcData{})},
		QUERY_3: lsquery{`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`, reflect.TypeOf(DataVw1{})},		
	}
}

type TBL_EnterpriseList struct {
	Enterprise string	
}

type DataVw1 struct {
	vip        string
	swVer      string
	enterprise string
	name       string
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
