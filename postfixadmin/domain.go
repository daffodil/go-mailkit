

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

//= A domain is a database in postfix
type Domain struct {
	Domain string 		` json:"domain"`
	Description string 	` json:"description"`
	Aliases int 		` json:"aliases"`
	Mailboxes int 		` json:"mailboxes"`
	MaxQuota int 		` json:"maxquota"`
	Quota int 			` json:"quota"`
	Transport string	` json:"transport"`
	BackupMx int 		` json:"backupmx"`
	Created string		` json:"created"`
	Modified string		` json:"modified"`
	Active int 			` json:"active"`
}

func(me Domain) TableName() string {
	return TableNames["domain"]
}

//= Ajax struct for `domain`
type DomainPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Error string `json:"error"`
}


//= Gets a domain from db, or its error
func GetDomain(domain_name string) (Domain, error) {
	var dom Domain
	var err error
	Dbo.Where("domain = ? ", domain_name).Order("domain").Find(&dom)
	return dom, err
}

//=  /ajax/domain/{domain}
func DomainAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("DomainAjaxHandler")
	vars := mux.Vars(req)

	payload := DomainPayload{}
	payload.Success = true

	var err error
	payload.Domain, err = GetDomain(vars["domain"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
