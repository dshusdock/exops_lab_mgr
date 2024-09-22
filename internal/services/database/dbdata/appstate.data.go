package dbdata

import (
	d "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"reflect"
	"time"
)

type AppState struct {
	Id 	   int
	CreateTime string
	State      string
}

var APP_STATE_VIEWS = make (map[string]viewMap)

func init() {
	APP_STATE_VIEWS["VIEW_ALL"] = viewMap{"select * from AppState", reflect.TypeOf(AppState{})}
	
}



func SetAppState() {
	dt := time.Now()
	nt := dt.Format("2006-01-02 15:04:05")
	str := fmt.Sprintf(`INSERT INTO AppState (create_time, state) VALUES ("%s", "active")`, nt)
	fmt.Println(str)
	d.WriteLocalDB(str)
}

func GetAppState() string {
	rows := d.ReadLocalDB("SELECT state FROM AppState")
	var state string
	for rows.Next() {
		rows.Scan(&state)
	}
	return state
}



