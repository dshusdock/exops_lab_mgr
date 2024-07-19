package tablevw

import (
	"dshusdock/tw_prac1/config"
	con "dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

type TableDef struct {
	Table       string
	Header      []string
	Tbl         []con.RowData
	TblSlice    []con.RowData
	SrchSlice   []con.RowData
	MaxRows     int
	RowCnt      int
	Start       int
	End         int
	Query       string
	SearchInput string
}

type TableVw struct {
	App        *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       TableDef
	Htmx       any
}

var AppTableVw *TableVw



func init() {
	AppTableVw = &TableVw{
		Id:         "tablevw",
		RenderFile: "table-view",
		ViewFlags:  []bool{false, true},
		Data: TableDef{
			Table:       "",
			Header:      nil,
			Tbl:         nil,
			TblSlice:    nil,
			SrchSlice:   nil,
			MaxRows:     10,
			RowCnt:      0,
			Start:       0,
			End:         0,
			Query:       "",
			SearchInput: "",
		},
		Htmx: nil,
	}

	
}



func (m *TableVw) RegisterView(app *config.AppConfig) *TableVw {
	AppTableVw.App = app
	return AppTableVw
}

func (m *TableVw) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[tablevw] In ProcessRequest")

	slog.Info("[tablevw] Entering Process Request")
	s := d.Get("event")
	fmt.Println("event is:" + s)

	switch s {
	case con.EVENT_CLICK:
		// m.processClickEvent(w, d)
	case con.EVENT_SEARCH:
		// m.processSearchEvent(w, d)
	}
}

func (m *TableVw) DisplaySQLTable(t string) {
	fmt.Println("HERE: " + t)
	m.ViewFlags[0] = true
	ptr := database.DB_VIEW_TYPE_MAP[t]
	fmt.Println("header: ", ptr.Header[0])
	fmt.Println("header: ", ptr.Header[1])
	fmt.Println("header: ", ptr.Header[2])
	fmt.Println("header: ", ptr.Header[3])

	// m.Data.Start = 0
	// m.Data.Tbl = m.readTableData(t)
	// m.Data.Table = t

	// var end int
	// m.Data.RowCnt = len(m.Data.Tbl)

	// if m.Data.RowCnt > m.Data.MaxRows {
	// 	end = m.Data.MaxRows
	// } else {
	// 	end = m.Data.RowCnt
	// }

	// m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]
	m.Data.Header = ptr.Header
}


// func (m *TableVw) DisplaySQLTable(t string) {
// 	fmt.Println("HERE")
// 	m.ViewFlags[0] = true
// 	ptr := models.BTN_SQL_QUERY_MAP[t]
// 	m.Data.Start = 0
// 	m.Data.Tbl = m.readTableData(t)
// 	m.Data.Table = t

// 	var end int
// 	m.Data.RowCnt = len(m.Data.Tbl)

// 	if m.Data.RowCnt > m.Data.MaxRows {
// 		end = m.Data.MaxRows
// 	} else {
// 		end = m.Data.RowCnt
// 	}

// 	m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]
// 	m.Data.Header = ptr.Header
// }

// func (m *TableVw) DisplayQueryTable(t string) {
// 	fmt.Println("HERE")
// 	m.ViewFlags[0] = true
// 	m.Data.Start = 0
// 	m.Data.Tbl, m.Data.Header = apis.ReadDBX(t)
// 	m.Data.Table = "Query Result"

// 	var end int
// 	m.Data.RowCnt = len(m.Data.Tbl)

// 	if m.Data.RowCnt > m.Data.MaxRows {
// 		end = m.Data.MaxRows
// 	} else {
// 		end = m.Data.RowCnt
// 	}

// 	m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]

// }

// func (m *TableVw) processSearchEvent(w http.ResponseWriter, d url.Values) {
// 	var rd []con.RowData
// 	fmt.Println("[tablevw] In processSearchEvent")
// 	fmt.Println("url.Values:" + d.Get("search"))
// 	key := d.Get("search")
// 	m.Data.SearchInput = key

// 	if key == "" {
// 		fmt.Println("Key is null")
// 		rd = m.Data.Tbl
// 	} else {
// 		for x := 0; x < m.Data.RowCnt; x++ {
// 			var row = m.Data.Tbl[x]
// 			if strings.Contains(strings.Join(row.Data, " "), key) {
// 				fmt.Println("got Row", row)
// 				rd = append(rd, row)
// 			}
// 		}
// 	}

// 	ln := len(rd)
// 	m.Data.RowCnt = ln
// 	m.Data.SrchSlice = rd
// 	if ln == 0 {
// 		m.Data.TblSlice = rd
// 	} else {
// 		var end int
// 		if ln > m.Data.MaxRows {
// 			end = m.Data.MaxRows
// 		} else {
// 			end = ln
// 		}
// 		m.Data.TblSlice = rd[m.Data.Start:end]
// 	}

// 	render.RenderMain(w, nil, m.App)
// }

// func (m *TableVw) processClickEvent(w http.ResponseWriter, d url.Values) {

// 	fmt.Println("[tablevw] In processClickEvent")
// 	lbl := d.Get("label")
// 	fmt.Println("Label: ", lbl)

// 	var end int
// 	switch d.Get("type") {
// 	case "button":
// 		if lbl == "Next" {
// 			fmt.Println("Next")

// 			if (m.Data.Start + (2 * m.Data.MaxRows)) < m.Data.RowCnt {
// 				m.Data.Start += m.Data.MaxRows
// 				end = m.Data.Start + m.Data.MaxRows

// 			} else {
// 				m.Data.Start += m.Data.MaxRows
// 				end = m.Data.Start + (m.Data.RowCnt - m.Data.Start)
// 			}
// 		}

// 		if lbl == "Previous" {
// 			fmt.Println("Previous")

// 			if (m.Data.Start - m.Data.MaxRows) > 0 {
// 				m.Data.Start -= m.Data.MaxRows
// 				end = m.Data.Start + m.Data.MaxRows

// 			} else {
// 				m.Data.Start = 0
// 				end = m.Data.MaxRows
// 			}
// 		}
// 	default:
// 	}
// 	if m.Data.SearchInput == "" {
// 		m.Data.TblSlice = m.Data.Tbl[m.Data.Start:end]
// 	} else {
// 		m.Data.TblSlice = m.Data.SrchSlice[m.Data.Start:end]
// 	}

// 	render.RenderMain(w, nil, m.App)
// }

// func (m *TableVw) readTableData(t string) []con.RowData {
// 	ptr := models.BTN_SQL_QUERY_MAP[t]
// 	pb := sidenav.SYS_SUB_BTN_LBL()
// 	rd := []con.RowData{}

// 	switch t {
// 	case pb.ENTERPISE_INFO:
// 		rd = apis.ReadDB[apis.EnterpriseInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.ENTERPISE_INFO
// 	case pb.ZONE_INFO:
// 		rd = apis.ReadDB[apis.ZoneInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.ZONE_INFO
// 	case pb.CCM_INFO:
// 		rd = apis.ReadDB[apis.CCMInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.CCM_INFO
// 	case pb.MEDIA_MGR_INFO:
// 		rd = apis.ReadDB[apis.MMInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.MEDIA_MGR_INFO
// 	case pb.MEDIA_GWY_INFO:
// 		rd = apis.ReadDB[apis.MGInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.MEDIA_GWY_INFO
// 	case pb.DEVICE_ZONE_INFO:
// 		rd = apis.ReadDB[apis.DeviceZoneInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.DEVICE_ZONE_INFO
// 	case pb.IQMAX_TURRET_INVENTORY:
// 		rd = apis.ReadDB[apis.IQMAXDeviceTurretInventory](ptr.SQL_STR)
// 		m.Data.Table = pb.IQMAX_TURRET_INVENTORY
// 	case pb.TURRET_INFO:
// 		rd = apis.ReadDB[apis.TurretInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.TURRET_INFO
// 	case pb.JOB_DETAILS_INFO:
// 		rd = apis.ReadDB[apis.JobDetailsInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.JOB_DETAILS_INFO
// 	case pb.CDI_COUNTS:
// 		rd = apis.ReadDB[apis.CdiCounts](ptr.SQL_STR)
// 		m.Data.Table = pb.CDI_COUNTS
// 	case pb.LICENSE_INFO:
// 		rd = apis.ReadDB[apis.LicenseInfo](ptr.SQL_STR)
// 		m.Data.Table = pb.LICENSE_INFO
// 	}

// 	return rd
// }

func (m *TableVw) ClearTable() {
	m.Data.Tbl = nil
	m.Data.TblSlice = nil
	m.Data.Table = ""
	m.Data.Header = nil
}
