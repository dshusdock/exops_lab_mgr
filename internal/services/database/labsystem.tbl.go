package database

import "fmt"

type DataVw1 struct {
	vip        string
	swVer      string
	enterprise string
	name       string
}

type MdcData struct {
	cab string
	cabULocaltion int16
	iso string
	name string
	serialNbr string
	iPAddress string
	vip string
	idracIp string
	swVer string
	serverType string
	enterprise string
	role string
	comments string
	vmLabServerHostIp string
}



func GetTableData() ([]DataVw1, error) {
	// An albums slice to hold data from returned rows.
	var data []DataVw1

	rows := DBA.Read(TBL_LAB_SYSTEM_QRY().QUERY_3)

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


type  labsystem_queries struct {
	QUERY_1		string
	QUERY_2		string
	QUERY_3		string
	QUERY_4		string
	QUERY_5		string
}

func TBL_LAB_SYSTEM_QRY() *labsystem_queries {
	return &labsystem_queries{
		QUERY_1:       	"select unique enterprise from LabSystem",
		QUERY_2:		"select * from LabSystem where enterprise = ?",
		QUERY_3: 		`SELECT vip, swVer, enterprise, name FROM LabSystem where role = "Unigy"`,
	}
}