package dbdata

import (
	"reflect"
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"

)

type UnigyDatabaseTargets struct {
	Enterprise string
	TargetIP string
	Status string
}

var UNIGY_DATABASE_TARGETS_VIEWS = make (map[string]viewMap)

func init() {
	UNIGY_DATABASE_TARGETS_VIEWS["VIEW_ALL"] = viewMap{"select * from UnigyDatabaseTargets", reflect.TypeOf(UnigyDatabaseTargets{})}
}

func (m *UnigyDatabaseTargets) GetAll() ([]con.RowData, error){ 
	rslt, err := d.ReadDBwithType[UnigyDatabaseTargets](UNIGY_DATABASE_TARGETS_VIEWS[VIEW_ALL].View)
	if err != nil {
		return nil, err
	}
	return rslt, nil  	
}

func (m *UnigyDatabaseTargets) GetView(qry string, parms ...string) ([]con.RowData, error){ return nil, nil }

func (m *UnigyDatabaseTargets) GetFieldList(fld string) ([]con.RowData, error){ return nil, nil }