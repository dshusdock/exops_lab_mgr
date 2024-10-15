package labsystemvw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	db "dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	b "dshusdock/tw_prac1/internal/views/base"
	"fmt"
	"log"
	"net/http"
)

type LSTableVW struct {
	App        *config.AppConfig
	// Id         string
	// RenderFile string
	// ViewFlags  []bool
	// LSTableVWData       TableDef
	// Htmx       any
}

var AppLSTableVW *LSTableVW

func init() {
	AppLSTableVW = &LSTableVW{
		App: nil,
	}
	messagebus.GetBus().Subscribe("Event:Click", AppLSTableVW.HandleMBusRequest)
}

func (m *LSTableVW) RegisterView(app *config.AppConfig) *LSTableVW {
	log.Println("Registering AppLSTableVW...")
	AppLSTableVW.App = app
	return AppLSTableVW
}

func (m *LSTableVW) RegisterHandler() constants.ViewHandler {
	return &LSTableVW{}
}

func (m *LSTableVW) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[LSTableVW] - Processing request")
	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)

	CreateLSTableVWData().ProcessHttpRequest(w, r)

}

func (m *LSTableVW) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[LSTableVW] - Processing MBus request")

	CreateLSTableVWData().ProcessMBusRequest(w, r)
}

func (m *LSTableVW) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[LSTableVW] - HandleRequest")
	obj := CreateLSTableVWData().ProcessHttpRequest(w, r)

	return obj
}
 

///////////////////// LSTable View LSTableVWData //////////////////////

type LSTableVWData struct {
	Base 			b.BaseTemplateparams
	LSTableVWData   TableDef
	SNData    		any
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

type AppLSTableVWData struct {
	Lbl string
}

func CreateLSTableVWData() *LSTableVWData {
	return &LSTableVWData{
		Base: b.GetBaseTemplateObj(),
		LSTableVWData: TableDef{
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
		SNData: nil,
	}
}

func (m *LSTableVWData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *LSTableVWData{	
	d := r.PostForm
	s := d.Get("label")
	
	switch s {
	case "upload":
		// render.RenderTemplate_new(w, nil, nil, con.RM_UPLOAD_MODAL)
		// renderview.RenderViewSvc.RenderTemplate(w, nil, con.RM_UPLOAD_MODAL)
	case "Table":
		m.Base.MainTable = true		
		m.Base.Cards = false
		m.LoadTableData(s)
		m.SNData = nil
	}
	return m
}

func (m *LSTableVWData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[%s] - Processing Internal request\n", "LSTableVWData")
	d := r.PostForm
	lbl := d.Get("label")
	
	switch lbl {
	case "Table":
		m.Base.MainTable = true		
		m.Base.Cards = false
		m.LoadTableData(lbl)
		m.SNData = nil
	}
	// render.RenderTemplate_new(w, nil, m, con.RM_TABLE)
	// renderview.RenderViewSvc.RenderTemplate(w, r, con.RM_TABLE)
}

func (m *LSTableVWData) LoadTableData(t string) error{
	var err error
	fmt.Println("\nDisplaying SQL Table: ", t)
	// m.ViewFlags[0] = true
	ptr := q.DB_VIEW_TYPE_MAP[t]

	m.LSTableVWData.Start = 0
	// m.LSTableVWData.Tbl, err = db.ReadTableData(t)
	m.LSTableVWData.Tbl, err = dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetAll()
	if err != nil {
		fmt.Println("Error in LoadTableData: ", err)
		return err
	}
	m.LSTableVWData.Table = t

	var end int
	m.LSTableVWData.RowCnt = len(m.LSTableVWData.Tbl)

	if m.LSTableVWData.RowCnt > m.LSTableVWData.MaxRows {
		end = m.LSTableVWData.MaxRows
	} else {
		end = m.LSTableVWData.RowCnt
	}

	m.LSTableVWData.TblSlice = m.LSTableVWData.Tbl[m.LSTableVWData.Start:end]
	m.LSTableVWData.HdrDef = ptr.HdrDef
	return nil
}

func (m *LSTableVWData) LoadTblDataByQuery(qry string) error{
	var err error
	fmt.Println("\nLoadDataByQuery - ", qry)
	// m.ViewFlags[0] = true
	ptr := q.DB_VIEW_TYPE_MAP["Table"]

	m.LSTableVWData.Start = 0
	m.LSTableVWData.Tbl, err = db.ReadTblWithQry(qry)
	if err != nil {
		fmt.Println("Error in LoadTblDataByQuery: ", err)
		return err
	}
	m.LSTableVWData.Table = "QUERY"

	var end int
	m.LSTableVWData.RowCnt = len(m.LSTableVWData.Tbl)

	if m.LSTableVWData.RowCnt > m.LSTableVWData.MaxRows {
		end = m.LSTableVWData.MaxRows
	} else {
		end = m.LSTableVWData.RowCnt
	}

	m.LSTableVWData.TblSlice = m.LSTableVWData.Tbl[m.LSTableVWData.Start:end]
	m.LSTableVWData.HdrDef = ptr.HdrDef
	return nil
}
