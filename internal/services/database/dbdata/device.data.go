package dbdata

import (
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"reflect"
)

type DeviceIfc struct {
	Ready bool
}

type Device struct {
	ID         int        
	Enterprise string   
	DeviceType string 
	Equipped   string
	Location   string 
	IP         string    
	MAC       string    	  
	ParentZone string
}

type viewObj1 struct {
	Enterprise string
	Ip 	   string
	Active_zone int
}

var DEVICE_VIEWS = make (map[string]viewMap)

func init() {
	DEVICE_VIEWS["VIEW_ALL"] = viewMap{"select * from Device", reflect.TypeOf(Device{})}
	DEVICE_VIEWS["VIEW_1"] = viewMap{`select * from Device where type = "%s" `, reflect.TypeOf(Device{})}
}

func (m *DeviceIfc) GetAll() ([]con.RowData, error) {
	rslt, err := d.ReadDBwithType[Device](DEVICE_VIEWS[VIEW_ALL].View)
	if err != nil {
		return nil, err
	}
	return rslt, nil  	
}

func (m *DeviceIfc) GetView(qry string, parms ...string) ([]con.RowData, error){
	var str string
	fmt.Println("In Device GetView: ", qry, len(parms))
	switch qry {
	case VIEW_1:
		str = fmt.Sprintf(DEVICE_VIEWS[VIEW_1].View, parms[0])
		fmt.Println("In Device GetView: ", str)
	}
	rslt, err := d.ReadDBwithType[Device](str)
	if err != nil {
		return nil, err
	}
	
	return rslt, nil
}

func (m *DeviceIfc) GetFieldList(fld string) ([]con.RowData, error){
	return []con.RowData{}, nil
}
