package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
)

type DBAccess struct {
	DBHandle *sql.DB
	active   bool
}

var DBA *DBAccess

func (m *DBAccess) Init() {
	var dba = DBAccess{
		DBHandle: nil,
	}

	dba.DBHandle = apis.Connect()
	DBA = &dba
	dba.active = true
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


