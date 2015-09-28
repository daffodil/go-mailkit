

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"sync"
)



// A memory cache for domains
var domainsMap map[string]Domain

// Load/Reload the domains
func LoadDomainsCache() error {

	domainsMap = make(map[string]Domain)
	domains, err := LoadDomains()
	if err != nil {
		return err
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	for _, dom := range domains {
		domainsMap[dom.Domain] = dom
	}
	mutex.Unlock()
	fmt.Println(domainsMap)
	return nil
}


// Load domains from database
func LoadDomains() ([]Domain, error) {
	var rows []Domain
	var err error
	Dbo.Where("domain <> ?", "ALL").Find(&rows)
	return rows, err
}

// Check a domain exists
// TODO: decide if active
func DomainExists(domain string) bool {

	if domainsMap == nil {
		LoadDomainsCache()
	}

	_, ok := domainsMap[domain]
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



// Handles /ajax/domains
func AjaxHandlerDomains(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("DomainsAjaxHandler")

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
