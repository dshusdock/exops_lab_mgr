package account_mgmt

import (
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/local"
	"errors"
	"fmt"
	"log"
)

type AccountMgrSvc struct {
	data string
}

var AcctMgrSvc *AccountMgrSvc

func init() {
	AcctMgrSvc = &AccountMgrSvc{ data: "Account Manager Service"} 
	log.Println("Initializing account manager service")
}

func CreateAccount(ac con.AccountInfo) error {

	un := local.GetUserNames() 
	for _, u := range un {
		if u.Data[0] == ac.Username {
			fmt.Println("Username already exists")
			return errors.New("Username already exists")
		}
	}
	
	str := fmt.Sprintf(`INSERT into User (firstname, lastname, email, username, password) values("%s","%s","%s","%s","%s")`, 
	ac.FirstName, ac.LastName, ac.Email, ac.Username, ac.Password)

	d.WriteLocalDB(str)
	return nil
}


