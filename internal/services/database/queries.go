package database

type  labsystem_queries struct {
	ENTERPISE		string
}

func SYS_SUB_BTN_LBL() *labsystem_queries {
	return &labsystem_queries{
		ENTERPISE:         "select unique enterprise from LabSystem",
	}
}