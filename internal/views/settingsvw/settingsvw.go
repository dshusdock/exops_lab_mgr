package settingsvw

import (
	"crypto/tls"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	"dshusdock/tw_prac1/internal/services/session"
	"dshusdock/tw_prac1/internal/services/unigy/unigydata"
	"dshusdock/tw_prac1/internal/views/base"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type SettingsVw struct {
	App *config.AppConfig

}

var AppSettingsVw *SettingsVw

func init() {
	AppSettingsVw = &SettingsVw{
		App: nil,
	}
	gob.Register(SettingsVwData{})
}

func (m *SettingsVw) RegisterView(app *config.AppConfig) *SettingsVw{
	log.Println("Registering AppSettingsVw...")
	AppSettingsVw.App = app
	return AppSettingsVw
}

func (m *SettingsVw) RegisterHandler() constants.ViewHandler {
	return &SettingsVw{}
}

func (m *SettingsVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	slog.Info("Processing request", "ID", "HeaderVw")
}

func (m *SettingsVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{ 

	return nil
}

func (m *SettingsVw) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[SettingsVw] - HandleRequest")
	var obj SettingsVwData

	if session.SessionSvc.SessionMgr.Exists(r.Context(), "settingsvw") {
		obj = session.SessionSvc.SessionMgr.Pop(r.Context(), "settingsvw").(SettingsVwData)
	} else {
		obj = *CreateSettingsVwData()	
	}
	obj.ProcessHttpRequest(w, r)	
	session.SessionSvc.SessionMgr.Put(r.Context(), "settingsvw", obj)

	return obj
}

///////////////////// Settings View Data //////////////////////

type SettingsVwData struct {
	Base base.BaseTemplateparams
	Data any
	Lbl string
	View int
	LastSynchTIme string
}

func CreateSettingsVwData() *SettingsVwData {
	return &SettingsVwData{
		Base: base.GetBaseTemplateObj(),
		Data: nil,
		Lbl: "Settings",
		View: constants.RM_SETTINGS_MODAL,
		LastSynchTIme: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (m *SettingsVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *SettingsVwData{
	fmt.Println("[settingsvw] - Processing request")

	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		m.View = constants.RM_UPLOAD_MODAL
		// render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "Test Button":
		fmt.Println("Test Button Clicked ok")		
	
	case "Test Button2":
		fmt.Println("Test Button2 Clicked")
		// database.PrintTableData()	
		// cardsvw.AppCardsVW.LoadCardData()
		unigydata.PopulateZoneInfoTable()
		// unigystatus.GetServerState("10.205.176.115")
	case "Test Button3":
		fmt.Println("Test Button3 Clicked")
		// token.TestEncrypt2()
		unigydata.PopulateDeviceTableByEnterprise("Dopey")
	case "Test Button4":
		fmt.Println("Test Button4 Clicked")
		
	case "Enter Button":
		fmt.Println("Enter Button Clicked")
		s := d.Get("ip")
		fmt.Println("IP: ", s)

		database.ConnectUnigyDB(s)
	case "Close Button":
		fmt.Println("Close Button Clicked")
		s := d.Get("ip")
		fmt.Println("IP: ", s)

		database.CloseUnigyDB()
	case "Zone Data Synch":
		fmt.Println("Zone Data Synch Clicked")		
		unigydata.PopulateZoneInfoTable()
	case "Target Synch":
		fmt.Println("Zone Data Synch Clicked")		
		unigydata.IdentifyValidDbEndpoints()
	case "Device Synch":
		fmt.Println("Device Synch Clicked")
		ent, _ := dbdata.GetDBAccess(dbdata.LAB_SYSTEM).GetFieldList("enterprise")
		for _, el := range ent {
			unigydata.PopulateDeviceTableByEnterprise(el.Data[0])
		}				
	}
	return m
}



func testHttp() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://10.205.176.114/haservices/CheckHAStatus")
	if err != nil {
		fmt.Println("Got an error" + err.Error())		
	}

	tasks, err := io.ReadAll(resp.Body)
    if err != nil {
        
    }
	fmt.Println(string(tasks))
}

func htmlParser() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://10.205.176.114/haservices/CheckHAStatus")
	if err != nil {
		fmt.Println("Got an error" + err.Error())					
	}

	 // Use the html package to parse the response body from the request
	 doc, err := html.Parse(resp.Body)
	 if err != nil {
		 fmt.Println("Error:", err)
		 return
	 }

	 // Function to recursively traverse the HTML node tree
	 var traverse func(*html.Node)
	 traverse = func(n *html.Node) {
		 if n.Type == html.TextNode {
			 fmt.Println("-->" + n.Data) // Print the name of the HTML element
		 }
		 for c := n.FirstChild; c != nil; c = c.NextSibling {
			 traverse(c)
		 }
	 }
 
	 // Traverse the HTML document
	 traverse(doc)
}



/*
func (c *APIClient) GetTasks() ([]byte, error) {
    conf := config.GetInstance()
    url := fmt.Sprintf("%s/myurl", conf.GetConfig().APIUrl)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.WithError(err).Errorf("Error creating HTTP request")
        return nil, err
    }

    // Add headers
    req.Header.Add("Authorization", conf.GetConfig().APIToken)
    req.Header.Add("Accept", "application/json")

    log.Info("Retrieving tasks from the API")
    resp, err := c.client.Do(req)
    if err != nil {
        log.WithError(err).Errorf("Error retrieving tasks from the backend")
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        errMsg := fmt.Sprintf("Received status: %s", resp.Status)
        err = errors.New(errMsg)
        log.WithError(err).Error("Error retrieving tasks from the backend")
        return nil, err
    }

    tasks, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.WithError(err).Error("Error reading tasks response body")
        return nil, err
    }

    log.Info("The tasks were successfully retrieved")

    return tasks, nil
}
*/
	


