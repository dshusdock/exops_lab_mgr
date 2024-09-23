package tables

import (
	"fmt"
	"time"
	"dshusdock/tw_prac1/internal/services/database"
)

func SetAppState() {
	dt := time.Now()
	nt := dt.Format("2006-01-02 15:04:05")
	str := fmt.Sprintf(`INSERT INTO AppState (create_time, state) VALUES ("%s", "active")`, nt)
	fmt.Println(str)
	database.WriteLocalDB(str)
}

func GetAppState() string {
	rows, _ := database.ReadLocalDB("SELECT state FROM AppState")
	var state = "ERROR"
	for rows.Next() {
		rows.Scan(&state)
	}
	return state
}
