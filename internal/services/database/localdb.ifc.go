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

func init() {
	LocalDB = &LocalDBAccess{
		Name: "LocalDB",
		DBHandle: nil,
		active:   false,
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
	// defer apis.Close(LocalDB.DBHandle)
	return apis.Read(LocalDB.DBHandle, sql)
}

func ReadDBwithType[T any](sql string) ([]con.RowData, error) {	
	rslt, err := apis.ReadDB[T](LocalDB.DBHandle, sql)
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
	ptr := q.DB_VIEW_TYPE_MAP[t]
	pb := con.HDR_BTN_LBL()

	fmt.Println("\nReadTableData: ", t)
	switch t {
	case pb.HDR_BTN_TABLE:
		rd, err := apis.ReadDB[q.MdcData](LocalDB.DBHandle, ptr.SQL_STR)
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
	rslt, err := apis.ReadDB[q.MdcData](LocalDB.DBHandle, sql)
	if err != nil {
		fmt.Println("Error in ReadTblWithQry: ", err)
		return nil, err
	}
	return rslt, nil
}




