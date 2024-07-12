package apis

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB){
	var db *sql.DB

	// Capture connection properties.
	cfg := mysql.Config{
        User:   "root",//os.Getenv("DBUSER"),
        Passwd: "my-secret-pw",//os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testdb",
        AllowNativePasswords: true,
    }
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
    fmt.Println("Connected!")
    return db
}

func Disconnect(db *sql.DB) {
    db.Close()
}

func Write(db *sql.DB, sql string) {
    _, err := db.Exec(sql)
    if err != nil {
        log.Fatal(err)
    }
}

func Read(db *sql.DB, sql string) (*sql.Rows) {    
    rows, err := db.Query(sql)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    return rows
}
