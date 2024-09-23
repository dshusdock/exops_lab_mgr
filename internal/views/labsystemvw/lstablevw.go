package labsystemvw

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	db "dshusdock/tw_prac1/internal/services/database"

	// "dshusdock/tw_prac1/internal/services/database/dbdata"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type AppLSTableVWData struct {
	Lbl string
}



type TableDef struct {
	Table       string
	HdrDef      []con.HeaderDef
	Tbl         []con.RowData
	TblSlice    []con.RowData
	SrchSlice   []con.RowData
	MaxRows     int
	RowCnt      int
	Start       int
	End         int
	Query       string
	SearchInput string
	Width       []int
}

type LSTableVW struct {
	App        *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       TableDef
	Htmx       any
}

var AppLSTableVW *LSTableVW

func init() {
	AppLSTableVW = &LSTableVW{
		Id:         "lstablevw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: TableDef{
			Table:       "",
			HdrDef:      nil,
			Tbl:         nil,
			TblSlice:    nil,
			SrchSlice:   nil,
			MaxRows:     10,
			RowCnt:      0,
			Start:       0,
			End:         0,
			Query:       "",
			SearchInput: "",
			Width:       nil,
		},
	}

	messagebus.GetBus().Subscribe("Event:Click", AppLSTableVW.ProcessInternalRequest)
}

func (m *LSTableVW) RegisterView(app config.AppConfig) *LSTableVW {
	log.Println("Registering AppLSTableVW...")
	AppLSTableVW.App = &app
	return AppLSTableVW
}

func (m *LSTableVW) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[LSTableVW] - Processing request")
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, con.RM_UPLOAD_MODAL)
	}
}

func (m *LSTableVW) ProcessInternalRequest(w http.ResponseWriter, d url.Values) {

	fmt.Printf("[%s] - Processing Internal request\n", m.Id)	
	lbl := d.Get("label")
	
	switch lbl {
	case "Table":
		m.App.MainTable = true
		m.App.Cards = false
		m.LoadTableData(lbl)
	}
	render.RenderTemplate_new(w, nil, m.App, con.RM_TABLE)
}

func (m *LSTableVW) LoadTableData(t string) error{
	var err error
	fmt.Println("\nDisplaying SQL Table: ", t)
	m.ViewFlags[0] = true
	ptr := q.DB_VIEW_TYPE_MAP[t]

	m.Data.Start = 0
	// m.Data.Tbl, err = db.ReadTableData(t)
	m.Data.Tbl, err = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetAll()
	if err != nil {
		fmt.Println("Error in LoadTableData: ", err)
		return err
	}
	m.Data.Table = t

	var end int
	m.Data.RowCnt = len(m.Data.Tbl)

	if m.Data.RowCnt > m.Data.MaxRows {
		end = m.Data.MaxRows
	} else {
		end = m.Data.RowCnt
	}

	m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]
	m.Data.HdrDef = ptr.HdrDef
	return nil
}

func (m *LSTableVW) LoadTblDataByQuery(qry string) error{
	var err error
	fmt.Println("\nLoadDataByQuery - ", qry)
	m.ViewFlags[0] = true
	ptr := q.DB_VIEW_TYPE_MAP["Table"]

	m.Data.Start = 0
	m.Data.Tbl, err = db.ReadTblWithQry(qry)
	if err != nil {
		fmt.Println("Error in LoadTblDataByQuery: ", err)
		return err
	}
	m.Data.Table = "QUERY"

	var end int
	m.Data.RowCnt = len(m.Data.Tbl)

	if m.Data.RowCnt > m.Data.MaxRows {
		end = m.Data.MaxRows
	} else {
		end = m.Data.RowCnt
	}

	m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]
	m.Data.HdrDef = ptr.HdrDef
	return nil
}
