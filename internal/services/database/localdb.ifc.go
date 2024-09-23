package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
	con "dshusdock/tw_prac1/internal/constants"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type LocalDBAccess struct {
	Name    	string
	DBHandle 	*sql.DB
	active   	bool
}

var LocalDB *LocalDBAccess
var DB_CONFIG mysql.Config

func init() {
	LocalDB = &LocalDBAccess{
		Name: "LocalDB",
		DBHandle: nil,
		active:   false,
	}

	DB_CONFIG = mysql.Config{
		User:                 "root",         //os.Getenv("DBUSER"),
		Passwd:               "my-secret-pw", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "testdb",
		AllowNativePasswords: true,
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

func WriteLocalDB(sql string) error{
	dbh, err := apis.Connect(DB_CONFIG)
	if err != nil {
		fmt.Println("Error in WriteLocalDB: ", err)
		return err
	}
	defer apis.Close(dbh)
	// apis.Write(LocalDB.DBHandle, sql)
	apis.Write(LocalDB.DBHandle, sql)
	return nil
}

func ReadLocalDB(sql string) (*sql.Rows, error) {
	dbh, err := apis.Connect(DB_CONFIG)
	if err != nil {
		fmt.Println("Error in WriteLocalDB: ", err)
		return nil, err
	}
	defer apis.Close(dbh)
	// defer apis.Close(LocalDB.DBHandle)
	// return apis.Read(LocalDB.DBHandle, sql)
	return apis.Read(dbh, sql), nil
}

func ReadDBwithType[T any](sql string) ([]con.RowData, error) {	
	dbh, err := apis.Connect(DB_CONFIG)
	if err != nil {
		fmt.Println("Error in WriteLocalDB: ", err)
		return nil, err
	}
	defer apis.Close(dbh)
	// rslt, err := apis.ReadDB[T](LocalDB.DBHandle, sql)
	rslt, err := apis.ReadDB[T](dbh, sql)
	if err != nil {
		fmt.Println("Error in ReadDBwithType: ", err)	
		return nil, err
	}
	return rslt, nil 
}

func CloseLocalDB() {
	apis.Close(LocalDB.DBHandle)
}






// Utilies
func ReadTableData(t string) ([]con.RowData, error) {
	dbh, err := apis.Connect(DB_CONFIG)
	if err != nil {
		fmt.Println("Error in WriteLocalDB: ", err)
		return nil, err
	}
	defer apis.Close(dbh)

	ptr := q.DB_VIEW_TYPE_MAP[t]
	pb := con.HDR_BTN_LBL()

	fmt.Println("\nReadTableData: ", t)
	switch t {
	case pb.HDR_BTN_TABLE:
		// rd, err := apis.ReadDB[q.MdcData](LocalDB.DBHandle, ptr.SQL_STR)
		rd, err := apis.ReadDB[q.MdcData](dbh, ptr.SQL_STR)
		if err != nil {
			fmt.Println("Error in ReadTableData: ", err)
			return nil, err
		}
		// m.Data.Table = pb.HDR_BTN_TABLE
		return rd, nil
	}
	return nil, nil	
}

func ReadTblWithQry(sql string) ([]con.RowData, error) {	
	dbh, err := apis.Connect(DB_CONFIG)
	if err != nil {
		fmt.Println("Error in WriteLocalDB: ", err)
		return nil, err
	}
	defer apis.Close(dbh)

	// rslt, err := apis.ReadDB[q.MdcData](LocalDB.DBHandle, sql)
	rslt, err := apis.ReadDB[q.MdcData](dbh, sql)
	if err != nil {
		fmt.Println("Error in ReadTblWithQry: ", err)
		return nil, err
	}
	return rslt, nil
}




