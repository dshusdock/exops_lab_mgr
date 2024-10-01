package dbdata

import (
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	"fmt"

	// "log/slog"
	"reflect"
)


type LabSystemIfc struct {
	Ready bool
}

type LabSystem struct {
	Cab               string
	CabULocation      string
	Iso               string
	Name              string
	SerialNbr         string
	IPAddress         string
	Vip               string
	IdracIp           string
	SwVer             string
	ServerType        string
	Enterprise        string
	Role              string
	Comments          string
	VmLabServerHostIp string
}

type LocalZoneData struct {
	Id 			int
	Enterprise 	string
	Zid 		string
	Vip  		string
	Ccm1 		string
	Ccm2 		string
	Ccm1_state 	string
	Ccm2_state 	string
	Online 		bool
	Status 		string
}

type VIEW_OBJ1 struct {
	Enterprise string	
}

type VIEW_OBJ2 struct {
	ServerType string	
}

type VIEW_OBJ3 struct {
	SWVer string	
}

var LAB_SYSTEM_VIEWS = make (map[string]viewMap)
var HdrDef []con.HeaderDef

func init() {
	LAB_SYSTEM_VIEWS[VIEW_ALL] = viewMap{"select * from LabSystem", reflect.TypeOf(LabSystem{})}
	LAB_SYSTEM_VIEWS[VIEW_1] = viewMap{"select unique enterprise from LabSystem", reflect.TypeOf(VIEW_OBJ1{})}
	LAB_SYSTEM_VIEWS[VIEW_2] = viewMap{"select * from LabSystem where enterprise = ?", reflect.TypeOf(LabSystem{})}
	LAB_SYSTEM_VIEWS[VIEW_3] = viewMap{`select unique serverType from LabSystem where enterprise = `, reflect.TypeOf(VIEW_OBJ2{})}
	LAB_SYSTEM_VIEWS[VIEW_4] = viewMap{`select unique enterprise from LabSystem where role="Unigy"`, reflect.TypeOf(VIEW_OBJ1{})}
	LAB_SYSTEM_VIEWS[VIEW_5] = viewMap{`select unique swVer from LabSystem`, reflect.TypeOf(VIEW_OBJ3{})}
	LAB_SYSTEM_VIEWS[VIEW_6] = viewMap{`select * from ZoneInfo where enterprise= `, reflect.TypeOf(VIEW_OBJ2{})}
	LAB_SYSTEM_VIEWS[VIEW_7] = viewMap{`select unique swVer from LabSystem where role='Unigy'`, reflect.TypeOf(VIEW_OBJ3{})}
	LAB_SYSTEM_VIEWS[VIEW_8] = viewMap{`select unique swVer from LabSystem where enterprise = `, reflect.TypeOf(VIEW_OBJ3{})}

	

	HdrDef = []con.HeaderDef{
		{Header: "CAB", Width: "width: 60px"}, 
		{Header: "U", Width: "width: 60px"}, 
		{Header: "ISO", Width: "width: 60px"}, 
		{Header: "Name", Width: "width: 200px"}, 
		{Header: "Serial#/Service Tag", Width: "width: 100px"}, 
		
		{Header: "IP", Width: "width: 100px"}, 
		{Header: "VIP", Width: "width: 100px"}, 
		{Header: "iDracVIP", Width: "width: 100px"}, 
		{Header: "Software Ver", Width: "width: 100px"}, 
		{Header: "Server Type", Width: "width: 100px"}, 

		{Header: "Enterprise", Width: "width: 100px"}, 
		{Header: "Role", Width: "width: 100px"}, 
		{Header: "Action", Width: "width: 400px, word-break: break-all"}, 
		{Header: "VM Lab Server", Width: "width: 100px"}, 		
	}
}

func (m *LabSystemIfc) GetAll() ([]con.RowData, error) {
	rslt, err := d.ReadDBwithType[LabSystem](LAB_SYSTEM_VIEWS["VIEW_ALL"].View)
	if err != nil {
		return nil, err
	}
	return rslt, nil  	
}

func (m *LabSystemIfc) GetView(qry string, parms ...string) ([]con.RowData, error){
	var str string
	var rslt []con.RowData
	var err error

	fmt.Println("In GetView: ", qry, len(parms))
	switch qry {
	case VIEW_6:
		str = fmt.Sprintf(LAB_SYSTEM_VIEWS[VIEW_6].View + "\"%s\"", parms[0])
		rslt, err = d.ReadDBwithType[LocalZoneData](str)
	case VIEW_8:
		str = fmt.Sprintf(LAB_SYSTEM_VIEWS[VIEW_8].View + "\"%s\"", parms[0])
		rslt, err = d.ReadDBwithType[VIEW_OBJ3](str)
	}
	if err != nil {
		return nil, err
	}
	
	return rslt, nil
}


func (m *LabSystemIfc) GetFieldList(fld string) ([]con.RowData, error){
	var rslt []con.RowData
	var err error

	switch fld {
	case "enterprise":
		rslt, err = d.ReadDBwithType[VIEW_OBJ1](LAB_SYSTEM_VIEWS["VIEW_1"].View)		
	case "swversion":
		rslt, err = d.ReadDBwithType[VIEW_OBJ3](LAB_SYSTEM_VIEWS["VIEW_5"].View)
	case "swversion_unigy":
		rslt, err = d.ReadDBwithType[VIEW_OBJ3](LAB_SYSTEM_VIEWS["VIEW_7"].View)
	case "enterprise_unigy":
		rslt, err = d.ReadDBwithType[VIEW_OBJ1](LAB_SYSTEM_VIEWS["VIEW_4"].View)
	}
	if err != nil {
		return nil, err
	}
	return rslt, nil
}

// Helper functions
func GetAllData() ([]LabSystem, error) {
	var rsltAry []LabSystem
	da, err := d.ReadDBwithType[LabSystem](LAB_SYSTEM_VIEWS["VIEW_ALL"].View)		
	if err != nil {
		return nil, err
	}

	for _, el := range da {
		obj := LabSystem{}
		obj.Cab = el.Data[0]
		obj.CabULocation = el.Data[1]
		obj.Iso = el.Data[2]
		obj.Name = el.Data[3]
		obj.SerialNbr = el.Data[4]
		obj.IPAddress = el.Data[5]
		obj.Vip = el.Data[6]
		obj.IdracIp = el.Data[7]
		obj.SwVer = el.Data[8]
		obj.ServerType = el.Data[9]
		obj.Enterprise = el.Data[10]
		obj.Role = el.Data[11]
		obj.Comments = el.Data[12]
		obj.VmLabServerHostIp = el.Data[13]
		rsltAry = append(rsltAry, obj)
	}

	return rsltAry, nil
}








