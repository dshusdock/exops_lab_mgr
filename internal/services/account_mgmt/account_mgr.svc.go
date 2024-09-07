package account_mgmt

import "log"

type AccountMgrSvc struct {
	data string
}

var AcctMgrSvc *AccountMgrSvc

func init() {
	AcctMgrSvc = &AccountMgrSvc{ data: "Account Manager Service"} 
	log.Println("Initializing account manager service")
}

