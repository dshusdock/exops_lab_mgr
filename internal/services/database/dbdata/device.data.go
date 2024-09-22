package dbdata

import (
	"reflect"
	"time"
)

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
