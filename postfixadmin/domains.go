

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
)

var DomainsMap map[string]Domain


func LoadDomainsMap() error {

	fmt.Println("Load Domains Map")
	domains, err := GetDomains()
	if err != nil {
		return err
	}
	for _, dom := range domains {
		DomainsMap[dom.Domain] = dom
	}
	fmt.Println(DomainsMap)
	return nil
}

func DomainExists(domain string) bool {

	_, ok := DomainsMap[domain]
	if ok  {
		return true
	}
	return false

}


type DomainsPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domains []Domain `json:"domains"`
	Error string `json:"error"`
}



func NewDomainsPayload() DomainsPayload {
	t := DomainsPayload{}
	t.Success = true
	t.Domains = make([]Domain, 0)
	return t
}

// Gets `Domains` from database
// TODO filter by domain in source
func GetDomains() ([]Domain, error) {
	var rows []Domain
	var err error
	Dbo.Where("domain <> ?", "ALL").Find(&rows)
	return rows, err
}

// Handles /ajax/domains
func DomainsAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("DomainsAjaxHandler")

	payload := NewDomainsPayload()

	var err error
	payload.Domains, err = GetDomains()
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
