

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"sync"
)

var mutex = &sync.Mutex{}

// A memory cache for domains
var domainsMap map[string]Domain


func LoadDomainsCache() error {

	fmt.Println("Load Domains Map")
	domainsMap = make(map[string]Domain)
	domains, err := GetDomains()
	if err != nil {
		return err
	}
	mutex.Lock()
	for _, dom := range domains {
		domainsMap[dom.Domain] = dom
	}
	mutex.Unlock()
	fmt.Println(domainsMap)
	return nil
}

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
func AjaxHandlerDomains(resp http.ResponseWriter, req *http.Request) {
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
