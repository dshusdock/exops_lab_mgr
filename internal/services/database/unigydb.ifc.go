package database

import (
	"database/sql"
	"dshusdock/tw_prac1/internal/apis"
	con "dshusdock/tw_prac1/internal/constants"
	q "dshusdock/tw_prac1/internal/services/database/queries"
	"fmt"

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

func ConnectUnigyDB(ip string) error{
	ipStr := ip + ":3306"

	cfg := mysql.Config{
		User:                 "dunkin",    //os.Getenv("DBUSER"),
		Passwd:               "dunkin123", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 ipStr,
		DBName:               "dunkin",
		AllowNativePasswords: true,
	}

	var err error
	UnigyDB.DBHandle, err = apis.Connect(cfg)
	if err != nil {
		// log.Println(err)
		return err
	}
	return nil
}

func WriteUnigyDB(sql string) {
	apis.Write(UnigyDB.DBHandle, sql)
}

func ReadUnigyDBwithType[T any](sql string) []con.RowData {	
	return apis.ReadDB[T](UnigyDB.DBHandle, sql)
}

func ReadUnigyDB(sql string) *sql.Rows {
	return apis.Read(UnigyDB.DBHandle, sql)
}
func CloseUnigyDB() {
	apis.Close(UnigyDB.DBHandle)
}

func UpdateLocalZoneInfo() {
	s := fmt.Sprintf(q.SQL_QUERIES_UNIGY["QUERY_1"].Qry )
		// fmt.Println(s)
		da := ReadUnigyDBwithType[q.TBL_NZData](s)
		
		fmt.Println("NZData: ", da)
		for _, el := range da {
			WriteLocalDB(fmt.Sprintf(q.SQL_QUERIES_LOCAL["QUERY_8"].Qry, el.Data[0], el.Data[1], el.Data[2], el.Data[3]))
		}
		
		// DisConnectDB()
}
