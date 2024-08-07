package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
	con "dshusdock/tw_prac1/internal/constants"

	"github.com/go-sql-driver/mysql"
)

type DBAccess struct {
	DBHandle *sql.DB
	active   bool
}

var DBA DBAccess

func Init() {
	cfg := mysql.Config{
		User:                 "root",         //os.Getenv("DBUSER"),
		Passwd:               "my-secret-pw", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "testdb",
		AllowNativePasswords: true,
	}

	DBA.DBHandle = apis.Connect(cfg)
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

func ReadDatabase[T any](sql string) []con.RowData {	
	return apis.ReadDB[T](DBA.DBHandle, sql)
}

func ReadTblWithQry(sql string) []con.RowData {	
	return apis.ReadDB[MdcData](DBA.DBHandle, sql)
}




