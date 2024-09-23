package dbdata

import (
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"reflect"
	"time"
)

type AppStateIfc struct {
	Ready bool
}

type AppState struct {
	Id 	   int
	CreateTime string
	State      string
}

var APP_STATE_VIEWS = make (map[string]viewMap)

func init() {
	APP_STATE_VIEWS["VIEW_ALL"] = viewMap{"select * from AppState", reflect.TypeOf(AppState{})}
	
}

func (m *AppStateIfc) RunQuery(qry string, parms ...string) ([]con.RowData, error){
	rslt, err := d.ReadDBwithType[LabSystem](qry)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}

func (m *AppStateIfc) GetAll() ([]con.RowData, error){
	rslt, err := d.ReadDBwithType[AppState]("select * from AppState")
	if err != nil {
		return nil, err
	}
	return rslt, nil
}

func (m *AppStateIfc) GetFieldList(fld string) ([]con.RowData, error){
	return nil, nil
}

func SetAppState() {
	dt := time.Now()
	nt := dt.Format("2006-01-02 15:04:05")
	str := fmt.Sprintf(`INSERT INTO AppState (create_time, state) VALUES ("%s", "active")`, nt)
	fmt.Println(str)
	d.WriteLocalDB(str)
}

func GetAppState() string {
	rows, _ := d.ReadLocalDB("SELECT state FROM AppState")
	var state = "ERROR"
	for rows.Next() {
		rows.Scan(&state)
	}
	return state
}



