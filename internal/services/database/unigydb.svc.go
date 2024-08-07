package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"

	"github.com/go-sql-driver/mysql"
)

type UnigyDBAccess struct {
	Enterprise string
	DBHandle *sql.DB
	active   bool
}

var UnigyDB *UnigyDBAccess

func init() {
	UnigyDB = &UnigyDBAccess{
		Enterprise: "Unigy",
		DBHandle: nil,
		active:   false,
	}
}


func (m *UnigyDBAccess) ConnectDB(ip string) {
	ipStr := ip + ":3306"

	cfg := mysql.Config{
		User:                 "dunkin",         //os.Getenv("DBUSER"),
		Passwd:               "dunkin123", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 ipStr,
		DBName:               "dunkin",
		AllowNativePasswords: true,
	}
	m.DBHandle = apis.Connect(cfg)
}

func (m *UnigyDBAccess) DisConnectDB(ip string) {
	apis.Disconnect(m.DBHandle)
}


