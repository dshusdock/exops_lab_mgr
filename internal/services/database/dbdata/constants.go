package dbdata

import ()

const (
	VIEW_ALL = "VIEW_ALL"
	VIEW_1 = "VIEW_1"
	VIEW_2 = "VIEW_2"
	VIEW_3 = "VIEW_3"
	VIEW_4 = "VIEW_4"
	VIEW_5 = "VIEW_5"
	VIEW_6 = "VIEW_6"
)

type viewMap struct {
	View string
	Model interface{}
}
