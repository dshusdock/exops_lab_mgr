package dbdata

import (
	"fmt"
	"reflect"
	d "dshusdock/tw_prac1/internal/services/database"
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

func WriteZoneInfoData(z ZoneInfo) {			
	str := fmt.Sprintf(`INSERT into ZoneInfo (enterprise, zid, vip, ccm1, ccm2, ccm1_state, ccm2_state, online, status) values("%s","%s","%s","%s","%s","%s","%s","%v","%s")`, 
		z.Enterprise, z.Zid, z.Vip, z.Ccm1, z.Ccm2, z.Ccm1State, z.Ccm2State, z.Online, z.Status)

		d.WriteLocalDB(str)
}