package status

import (
	"crypto/tls"
	"dshusdock/tw_prac1/config"
	"io"
	"strings"

	// "dshusdock/tw_prac1/internal/services/datetime"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type AppSettingsVwData struct {
	Lbl string
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
	case "vip":
		fmt.Println("--------------------------------------------------")
		htmlParser()
	case "sometarget":

		
	}
}



func htmlParser() {
	// m := make(map[string]int)

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
			//  fmt.Println("-->" + n.Data) // Print the name of the HTML element
			if !strings.Contains(n.Data, "########") && !strings.Contains(n.Data, "=======") {
			 	res1 := strings.Split(n.Data, ":")
			 	fmt.Println(res1[0])
			 	fmt.Println(res1[1])
			}

			 

		 }
		 for c := n.FirstChild; c != nil; c = c.NextSibling {
			 traverse(c)
		 }
	 }
 
	 // Traverse the HTML document
	 traverse(doc)
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

