package dbdata

import (
	"dshusdock/tw_prac1/internal/constants"
	con "dshusdock/tw_prac1/internal/constants"
)

type DBDataService struct {

}

const (
	LAB_SYSTEM = iota
	APP_STATE
	DEVICE
)

type DBDataAccess interface {
	GetView(qry string, parms ...string) ([]constants.RowData, error)
	GetAll() ([]con.RowData, error)
	GetFieldList(fld string) ([]con.RowData, error)
}

func GetDBAccess(tbl int) DBDataAccess {

	switch tbl {
	case LAB_SYSTEM:
		return &LabSystemIfc{ Ready: true }
	case APP_STATE:
		return &AppStateIfc{ Ready: true }
	case DEVICE:
		return &DeviceIfc{ Ready: true }
	}

	return nil
}
