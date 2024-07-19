package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
	con "dshusdock/tw_prac1/internal/constants"
)

type DBAccess struct {
	DBHandle *sql.DB
	active   bool
}

var DBA DBAccess

func Init() {
	DBA.DBHandle = apis.Connect()
	DBA.active = true
}

func (m *DBAccess) Write(sql string) {
	apis.Write(DBA.DBHandle, sql)
}

func (m *DBAccess) Read(sql string) (*sql.Rows) {
	return apis.Read(DBA.DBHandle, sql)
}


func (m *DBAccess) Disconnect() {
	apis.Disconnect(DBA.DBHandle)
}

func ReadTableData(t string) []con.RowData {
	ptr := DB_VIEW_TYPE_MAP[t]
	pb := con.HDR_BTN_LBL()
	rd := []con.RowData{}

	switch t {
	case pb.HDR_BTN_TABLE:
		rd = apis.ReadDB[MdcData](DBA.DBHandle, ptr.SQL_STR)
		// m.Data.Table = pb.HDR_BTN_TABLE
	}

	return rd
}


