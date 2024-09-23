package tables

import (
	d "dshusdock/tw_prac1/internal/services/database"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"fmt"
)

func GetTableData() ([]q.DataVw1, error) {
	// An albums slice to hold data from returned rows.
	var data []q.DataVw1

	rows, _ := d.ReadLocalDB(q.SQL_QUERIES_LOCAL["QUERY_3"].Qry)

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dv q.DataVw1
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

