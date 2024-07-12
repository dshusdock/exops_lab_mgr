package database

import (
	
)

func SetAppState(s string) {
	DBA.Write("INSERT INTO appstate VALUES (s)")
}

func GetAppState() string {
	rows := DBA.Read("SELECT state FROM appstate")
	var state string
	for rows.Next() {
		rows.Scan(&state)
	}
	return state
}
