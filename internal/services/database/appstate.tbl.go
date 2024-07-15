package database

import (
	"fmt"
	"time"
)



func SetAppState() {
	dt := time.Now()
	nt := dt.Format("2006-01-02 15:04:05")
	str := fmt.Sprintf(`INSERT INTO AppState (create_time, state) VALUES ("%s", "active")`, nt)
	fmt.Println(str)
	DBA.Write(str)
}

func GetAppState() string {
	rows := DBA.Read("SELECT state FROM AppState")
	var state string
	for rows.Next() {
		rows.Scan(&state)
	}
	return state
}
