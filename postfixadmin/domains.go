

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
)





// Load domains from database
func LoadDomains() ([]Domain, error) {
	var rows []Domain
	var err error
	Dbo.Where("domain <> ?", "ALL").Find(&rows)
	return rows, err
}




type DomainsPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domains []Domain `json:"domains"`
	Error string `json:"error"`
}



// Handles /ajax/domains
func AjaxHandlerDomains(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("DomainsAjaxHandler")
	if AjaxAuth(resp, req) == false {
		return
	}

	payload := DomainsPayload{}
	payload.Success = true
	//t.Domains = make([]Domain, 0)

	var err error
	payload.Domains, err = LoadDomains()
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
