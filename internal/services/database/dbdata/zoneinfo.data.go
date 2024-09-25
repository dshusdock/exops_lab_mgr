package dbdata

import (
	"fmt"
	"reflect"
	d "dshusdock/tw_prac1/internal/services/database"
	con "dshusdock/tw_prac1/internal/constants"

)

type ZoneInfo struct {
	Id		  int
	Enterprise string
	Zid		  string
	Vip 	  string
	Ccm1	  string
	Ccm2	  string
	Ccm1State string
	Ccm2State string
	Online	  string
	Status	  string
}

var ZONE_INFO_VIEWS = make (map[string]viewMap)

func init() {
	ZONE_INFO_VIEWS["VIEW_ALL"] = viewMap{"select * from ZoneInfo", reflect.TypeOf(ZoneInfo{})}
}

func (m *ZoneInfo) GetAll() ([]con.RowData, error) {
	rslt, err := d.ReadDBwithType[ZoneInfo](ZONE_INFO_VIEWS[VIEW_ALL].View)
	if err != nil {
		return nil, err
	}
	return rslt, nil  	
}

func (m *ZoneInfo) GetView(qry string, parms ...string) ([]con.RowData, error) { return nil, nil }

func (m *ZoneInfo) GetFieldList(fld string) ([]con.RowData, error) { return nil, nil }

// Helper functions
func WriteZoneInfoData(z ZoneInfo) {			
	str := fmt.Sprintf(`INSERT into ZoneInfo (enterprise, zid, vip, ccm1, ccm2, ccm1_state, ccm2_state, online, status) values("%s","%s","%s","%s","%s","%s","%s","%v","%s")`, 
		z.Enterprise, z.Zid, z.Vip, z.Ccm1, z.Ccm2, z.Ccm1State, z.Ccm2State, z.Online, z.Status)

		d.WriteLocalDB(str)
}