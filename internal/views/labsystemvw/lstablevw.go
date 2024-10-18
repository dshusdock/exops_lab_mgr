package labsystemvw

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
	db "dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"dshusdock/tw_prac1/internal/services/messagebus"
	"dshusdock/tw_prac1/internal/services/session"
	b "dshusdock/tw_prac1/internal/views/base"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
)

type LSTableVW struct {
	App        *config.AppConfig
}

var AppLSTableVW *LSTableVW

func init() {
	AppLSTableVW = &LSTableVW{
		App: nil,
	}
	gob.Register(LSTableVWData{})
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

func (m *LSTableVW) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{
	fmt.Println("[LSTableVW] - Processing MBus request")

	obj := CreateLSTableVWData().ProcessMBusRequest(w, r)
	return obj
}

func (m *LSTableVW) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[LSTableVW] - HandleRequest")


	var obj LSTableVWData

	if session.SessionSvc.SessionMgr.Exists(r.Context(), "lstablevw") {
		obj = session.SessionSvc.SessionMgr.Pop(r.Context(), "lstablevw").(LSTableVWData)
	} else {
		obj = *CreateLSTableVWData()	
	}

	obj.ProcessHttpRequest(w, r)	
	session.SessionSvc.SessionMgr.Put(r.Context(), "lstablevw", obj)

	return obj
}
 

///////////////////// LSTable View LSTableVWData //////////////////////

type LSTableVWData struct {
	Base 			b.BaseTemplateparams
	LSTableVWData   TableDef
	View 			int
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
	e := d.Get("event")

	switch e {
	case con.EVENT_CLICK:
		m.ProcessClickEvent(w, r)
	case con.EVENT_SEARCH:
		// m.processSearchEvent(w, d)	
	
	default:
		m.ProcessRequest(w, r)
	}
	return m
}

func (m *LSTableVWData) ProcessClickEvent(w http.ResponseWriter, r *http.Request) *LSTableVWData{
	d := r.PostForm
	lbl := d.Get("label")
	str := d.Get("view_str")

	if lbl == "Table" {
		m.Base.MainTable = true		
		m.Base.Cards = false
		m.LoadTableData(lbl)
		m.SNData = nil
	} else {
		m.LoadTblDataByQuery(getListFromId(str, lbl))
		m.View = con.RM_TABLE_REFRESH
	}

	return m
}

func (m *LSTableVWData) ProcessRequest(w http.ResponseWriter, r *http.Request) *LSTableVWData{
	// d := r.PostForm
	// s := d.Get("label")

	
	return m
}

func (m *LSTableVWData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) *LSTableVWData{

	fmt.Printf("[%s] - Processing Internal request\n", "LSTableVWData")
	d := r.PostForm
	lbl := d.Get("label")
	id := d.Get("view_id")
	str := d.Get("view_str")
	fmt.Println("[LSTableVWData]label: ", lbl)
	fmt.Println("[LSTableVWData]view_id: ", id)
	fmt.Println("[LSTableVWData]view_str: ", str)
	
	switch lbl {
	case "Table":
		m.Base.MainTable = true		
		m.Base.Cards = false
		m.LoadTableData(lbl)
		m.SNData = nil
		return m
	}

	return m
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

func getListFromId(id string, lbl string) string {
	var str string
	part := "Select * from LabSystem where "

	switch id {
	case "enterprise":
		str = fmt.Sprintf(part + "Enterprise = \"%s\"", lbl)
	case "swver":
		str = fmt.Sprintf(part + "swVer = \"%s\"", lbl)
	case "Unigy":
		str = fmt.Sprintf(part + "Enterprise = \"%s\"", lbl)
	}
	fmt.Println("---->Query: ", str)
	return str
}

