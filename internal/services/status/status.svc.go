package status

import (
	"crypto/tls"
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/render"
	"strings"

	// "dshusdock/tw_prac1/internal/services/datetime"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type HAStatusServiceInfo struct {
	Name string
	Status string
	Remarks string
}

type HAStatusInfo struct {
	Enterprise string
	Server string
	ServerRole string
	ServerState	string
	NetworkState string
	OrchestratorState string
	ServerRemarks string
	ServerStarted string
	ServiceInfo	[]HAStatusServiceInfo
	TCPPort7800Status1 string
	TCPPort7800Status2 string
	FirewallStatus string
}

type StatusSvc struct {
	App *config.AppConfig
	Id         string
	RenderFile string
	ViewFlags  []bool
	Data       any
	Htmx       any
}

var AppStatusSvc *StatusSvc

func init() {
	AppStatusSvc = &StatusSvc{
		Id:         "statussvc",
		RenderFile: "",
		ViewFlags:  []bool{true},
		Data: "",
		Htmx: nil,
	}
}

func (m *StatusSvc) RegisterView(app config.AppConfig) *StatusSvc{
	log.Println("Registering AppStatusSvc...")
	AppStatusSvc.App = &app
	return AppStatusSvc
}

func (m *StatusSvc) ProcessRequest(w http.ResponseWriter, d url.Values) {
	fmt.Println("[statussvc] - Processing request")
	da := d.Get("data")
	tg := d.Get("target")

	fmt.Println("data: ", da)
	fmt.Println("target: ", tg)

	switch tg {
	case "ip":
		fmt.Println("--------------------------------------------------")
		// GetHaStatus(da)
		s :=getServerStatus(da)
		render.JSONResponse(w, s)
	case "sometarget":

		
	}
}

func getServerStatus(ip string) string{
	fmt.Println("Getting server status")
	fmt.Println("IP: ", ip)
	fmt.Println("--------------------------------------------------")
	
	rslt, err := GetHaStatus(ip)
	if err != nil {
		fmt.Println("Error getting the HA status")
		return "ERROR"
	}

	if ((rslt.ServerRole == "ACTIVE" || rslt.ServerRole == "STANDBY") && 
		rslt.ServerState == "AVAILABLE") {
		return "RUNNING"
	} else {
		return "DOWN"
	}
}

func GetHaStatus(vip string) (HAStatusInfo, error) {
	var retVal HAStatusInfo
	var info HAStatusServiceInfo
	retVal = HAStatusInfo{}

	rslt, err := CheckHAStatusParser(vip)
	if err != nil {
		return HAStatusInfo{}, err
	}

	if len(rslt) < 20 {
		return HAStatusInfo{}, fmt.Errorf("Error parsing the HA status")
	}

	for i, v := range rslt {

		switch i {
		case 0:
			retVal.Server = strings.Split(v, ":")[1]
			continue
		case 1:
			retVal.ServerRole = strings.Split(v, ":")[1]
			continue
		case 2:
			retVal.ServerState = strings.Split(v, ":")[1]
			continue
		case 3:
			retVal.NetworkState = strings.Split(v, ":")[1]
			continue
		case 4:
			retVal.OrchestratorState = strings.Split(v, ":")[1]
			continue
		case 5:
			retVal.ServerRemarks = strings.Split(v, ":")[1]
			continue
		case 6:
			retVal.ServerStarted = strings.Split(v, ":")[1]
			continue
		case 7:
			continue
		}

		if strings.Contains(v, "Service Name") {
			info = HAStatusServiceInfo{}
			info.Name = strings.Split(v, ":")[1]				
		}

		if strings.Contains(v, "Service Status") {
			info.Status = strings.Split(v, ":")[1]
		}

		if strings.Contains(v, "Remarks") {
			info.Remarks = strings.Split(v, ":")[1]
			retVal.ServiceInfo = append(retVal.ServiceInfo, info)
		}		
	}
	// fmt.Printf("%#v", retVal)
	return retVal, nil
}

func CheckHAStatusParser(ip string) ([]string, error) {
	var result []string

	// Accept the certificate
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	// str := "https://" + vip + "/haservices/CheckHAStatus"
	str := "https://" + ip + "/haservices/checkMyStatus"
	resp, err := client.Get(str)
	if err != nil {
		fmt.Println("Error getting the response from the URL")
		return nil, err
	}

	 // Use the html package to parse the response body from the request
	 doc, err := html.Parse(resp.Body)
	 if err != nil {
		fmt.Println("Error parsing the HTML document")
		return nil, err		
	 }

	 // Function to recursively traverse the HTML node tree
	 var traverse func(*html.Node)

	 traverse = func(n *html.Node) {
		 if n.Type == html.TextNode {

			if 	!strings.Contains(n.Data, "########") && 
			   	!strings.Contains(n.Data, "=======") && 
				!strings.Contains(n.Data, "HTTP 200 OK") {
				fmt.Println("===>" + n.Data)
				result = append(result, n.Data)
			}
		 }
		 for c := n.FirstChild; c != nil; c = c.NextSibling {
			 traverse(c)
		 }
	 }
 
	 // Traverse the HTML document
	 traverse(doc)
	 return result, nil
}




