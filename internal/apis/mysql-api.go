package apis

import (
	"database/sql"
	con "dshusdock/tw_prac1/internal/constants"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DBHandle2 *sql.DB = nil

func Connect(cfg mysql.Config) *sql.DB {
	var db *sql.DB

	cfg.Timeout, _ = time.ParseDuration("5s")
	
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("SQL Database Connected!...")
	return db
}

func Disconnect(db *sql.DB) {
	db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Connection Closed!")
	}
}

func Write(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func Read(db *sql.DB, sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func ReadDB[T any](db *sql.DB, s string) []con.RowData {
	var tableDef []T

	rows, err := db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e T

		s := reflect.ValueOf(&e).Elem()

		numCols := s.NumField()
		//fmt.Println("numCols is", numCols)
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		if err := rows.Scan(columns...); err != nil {
			log.Fatal(err)
		}
		tableDef = append(tableDef, e)
	}

	var rd = []con.RowData{}

	for i := 0; i < len(tableDef); i++ {
		values := reflect.ValueOf(tableDef[i])

		r := con.RowData{
			Data: nil,
		}
		for ii := 0; ii < values.NumField(); ii++ {
			f := values.Field(ii)
			r.Data = append(r.Data, checkReflect(f))
		}
		rd = append(rd, r)
	}

	return rd
}

// Check to see if this is a sql "null" type
func checkReflect(f reflect.Value) string {
	if f.Kind().String() == "struct" {
		val := f.Interface().(sql.NullString)

		if val.Valid {
			fmt.Println("Valid Data:", val.String)
			return val.String
		} else {
			return "null"
		}
	}
	return f.String()
}
