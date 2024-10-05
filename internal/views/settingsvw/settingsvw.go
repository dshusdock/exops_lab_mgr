package settingsvw

import (
	"crypto/tls"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"dshusdock/tw_prac1/internal/services/database"
	"dshusdock/tw_prac1/internal/services/database/dbdata"
	"dshusdock/tw_prac1/internal/services/unigy/unigydata"
	"time"

	// "dshusdock/tw_prac1/internal/services/unigy/unigystatus"

	// "dshusdock/tw_prac1/internal/views/cardsvw"
	"io"

	// "dshusdock/tw_prac1/internal/services/datetime"
	"fmt"
	"log"
	"net/http"

	// "net/url"

	"golang.org/x/net/html"
)

type AppSettingsVwData struct {
	Lbl string
}

type SettingsVw struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
	LastSynchTIme string
}

var AppSettingsVw *SettingsVw

func init() {
	AppSettingsVw = &SettingsVw{
		Id:         "settingsvw",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
		LastSynchTIme: time.Now().Format("2006-01-02 15:04:05"),
	}

}

func (m *SettingsVw) RegisterView(app config.AppConfig) *SettingsVw{
	log.Println("Registering AppSettingsVw...")
	AppSettingsVw.App = &app
	return AppSettingsVw
}

func (m *SettingsVw) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[settingsvw] - Processing request")

	val := m.App.SessionManager.Get(r.Context(), "LoggedIn")
	if val != true {
		// m.App.LoggedIn = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		// m.App.LoggedIn = true
	}

	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)

	switch s {
	case "upload":
		render.RenderTemplate_new(w, nil, nil, constants.RM_UPLOAD_MODAL)
	case "Test Button":
		fmt.Println("Test Button Clicked ok")		

		// datetime.Prac1()   
		// datetime.Prac2()
		// datetime.Prac3()

		// htmlParser()
		// unigydata.IdentifyValidDbEndpoints()
		m.App.SessionManager.Put(r.Context(), "btnpressed", "truedat")
	
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
		// token.TestDecrypt2()
		x := m.App.SessionManager.Get(r.Context(), "btnpressed")
		fmt.Println("Button Pressed: ", x)
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
	


