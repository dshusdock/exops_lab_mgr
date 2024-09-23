package dbdata

import (
	"reflect"
	"time"
	d "dshusdock/tw_prac1/internal/services/database"
	con "dshusdock/tw_prac1/internal/constants"
)

type DeviceIfc struct {
	Ready bool
}

type Device struct {
	ID         int        
	CreateTime time.Time 
	Enterprise string    
	IP         string    
	Type       string    
	Location   string   
}

var DEVICE_VIEWS = make (map[string]viewMap)

func init() {
	DEVICE_VIEWS["VIEW_ALL"] = viewMap{"select * from Device", reflect.TypeOf(Device{})}
}

func (m *DeviceIfc) RunQuery(qry string, parms ...string) ([]con.RowData, error){
	rslt, err := d.ReadDBwithType[LabSystem](qry)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}	
