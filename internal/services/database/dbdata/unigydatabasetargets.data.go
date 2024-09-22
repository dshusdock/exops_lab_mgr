package dbdata

import "reflect"

type UnigyDatabaseTargets struct {
	Enterprise string
	TargetIP string
	Status string
}

var UNIGY_DATABASE_TARGETS_VIEWS = make (map[string]viewMap)

func init() {
	UNIGY_DATABASE_TARGETS_VIEWS["VIEW_ALL"] = viewMap{"select * from UnigyDatabaseTargets", reflect.TypeOf(UnigyDatabaseTargets{})}
}
