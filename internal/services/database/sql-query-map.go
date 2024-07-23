package database

import (
	con "dshusdock/tw_prac1/internal/constants"
)

type SqlQueryDef struct {
	SQL_STR string
	HdrDef  []con.HeaderDef
}

// type SqlQueryMap struct {
// 	SqlQueryEntries []SqlQueryDef
// }

var DB_VIEW_TYPE_MAP map[string]SqlQueryDef

func init() {
	pa := con.HDR_BTN_LBL()

	DB_VIEW_TYPE_MAP = map[string]SqlQueryDef{
		pa.HDR_BTN_TABLE: {
			SQL_STR: "Select * from LabSystem",
			// Header:  []string{"CAB", "The U", "ISO", "Name", "Seial#/Service Tag", "IP", "VIP", "iDracVIP", "Software Ver", "Server Type", "Enterprise", "Role", "Action", "VM Lab Server"},
			HdrDef:  []con.HeaderDef{
				{Header: "CAB", Width: "width: 60px"}, 
				{Header: "U", Width: "width: 60px"}, 
				{Header: "ISO", Width: "width: 60px"}, 
				{Header: "Name", Width: "width: 200px"}, 
				{Header: "Seial#/ Service Tag", Width: "width: 100px"}, 
				
				{Header: "IP", Width: "width: 100px"}, 
				{Header: "VIP", Width: "width: 100px"}, 
				{Header: "iDracVIP", Width: "width: 100px"}, 
				{Header: "Software Ver", Width: "width: 100px"}, 
				{Header: "Server Type", Width: "width: 100px"}, 

				{Header: "Enterprise", Width: "width: 100px"}, 
				{Header: "Role", Width: "width: 100px"}, 
				{Header: "Action", Width: "width: 400px, word-break: break-all"}, 
				{Header: "VM Lab Server", Width: "width: 100px"}, 
				
			},			
		},
	}
}
