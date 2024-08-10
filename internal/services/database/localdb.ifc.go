package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
	con "dshusdock/tw_prac1/internal/constants"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type LocalDBAccess struct {
	Name    	string
	DBHandle 	*sql.DB
	active   	bool
}

var LocalDB *LocalDBAccess

func init() {
	LocalDB = &LocalDBAccess{
		Name: "LocalDB",
		DBHandle: nil,
		active:   false,
	}
	
	var err error
	if err != nil {
		log.Println(err)
	}
}

func ConnectLocalDB(ip string) error{
	ipStr := ip + ":3306"

	cfg := mysql.Config{
		User:                 "root",         //os.Getenv("DBUSER"),
		Passwd:               "my-secret-pw", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 ipStr,
		DBName:               "testdb",
		AllowNativePasswords: true,
	}

	var err error
	LocalDB.DBHandle, err = apis.Connect(cfg)
	if err != nil {
		// log.Println(err)
		return err
	}
	LocalDB.active = true
	return nil
}

func WriteLocalDB(sql string) {
	apis.Write(LocalDB.DBHandle, sql)
}

func ReadLocalDB(sql string) (*sql.Rows) {
	return apis.Read(LocalDB.DBHandle, sql)
}

func ReadLocalDBwithType[T any](sql string) []con.RowData {	
	return apis.ReadDB[T](LocalDB.DBHandle, sql)
}

func CloseLocalDB() {
	apis.Close(LocalDB.DBHandle)
}

// Utilies
func ReadTableData(t string) []con.RowData {
	ptr := q.DB_VIEW_TYPE_MAP[t]
	pb := con.HDR_BTN_LBL()
	rd := []con.RowData{}

	fmt.Println("\nReadTableData: ", t)
	switch t {
	case pb.HDR_BTN_TABLE:
		rd = apis.ReadDB[q.MdcData](LocalDB.DBHandle, ptr.SQL_STR)
		// m.Data.Table = pb.HDR_BTN_TABLE
	}

	return rd
}

func ReadTblWithQry(sql string) []con.RowData {	
	return apis.ReadDB[q.MdcData](LocalDB.DBHandle, sql)
}




