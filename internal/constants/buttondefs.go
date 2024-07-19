package constants

type header_btn_label struct {
	HDR_BTN_TABLE         string
}

func HDR_BTN_LBL() *header_btn_label {
	return &header_btn_label{
		HDR_BTN_TABLE:         "Table",
	}
}


// /////////////EXAMPLE///////////////
type side_nav_btn_lbl struct {
	SYSTEM          string
	USER            string
	RECORDING       string
	BUTTON          string
	RESOURCE_AOR    string
	OPEN_CONNECTION string
	LINE            string
	ZONE            string
}

func SIDE_NAV_BTN_LBL() *side_nav_btn_lbl {
	return &side_nav_btn_lbl{
		SYSTEM:          "System",
		USER:            "User",
		RECORDING:       "Recording",
		BUTTON:          "Button",
		RESOURCE_AOR:    "Resource AOR",
		OPEN_CONNECTION: "Open Connection",
		LINE:            "Line",
		ZONE:            "Zone",
	}
}



