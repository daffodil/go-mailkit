

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"errors"

	"github.com/cenkalti/log"
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

func IsDomainValid(domain_name string) error {

	if DomainExists(domain_name) == false {
		return errors.New("Domain `" + domain_name + "' does not exist")
	}
	return nil
}


//= Loads a domain row from db
//
// TODO check its in cache ?
func LoadDomain(domain_name string) (Domain, error) {
	var dom Domain
	var err error

	err = IsDomainValid(domain_name)
	if err != nil {
		return dom, err
	}
	Dbo.Where("domain = ? ", domain_name).Order("domain").Find(&dom)
	return dom, err
}


//= Ajax struct for `domain`
type DomainPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Error string `json:"error"`
}

//=  /ajax/domain/{domain}
func AjaxHandlerDomain(resp http.ResponseWriter, req *http.Request) {

	//fmt.Println("DomainAjaxHandler")
	log.Info("DomainAjaxHandler")

	vars := mux.Vars(req)

	payload := DomainPayload{}
	payload.Success = true

	var err error
	payload.Domain, err = LoadDomain(vars["domain"])
	if err != nil{
		log.Info(err.Error())
		payload.Error = "" + err.Error()
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}


// Ajax struct for `domain` all
type DomainAllPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Mailboxes []Mailbox `json:"mailboxes"`
	Aliases []Alias `json:"aliases"`
	Error string `json:"error"`
}


//  /ajax/domain/{domain}/all
func AjaxHandlerDomainAll(resp http.ResponseWriter, req *http.Request) {

	log.Info("DomainAllAjaxHandler")

	vars := mux.Vars(req)
	domain := vars["domain"]

	payload := DomainAllPayload{}
	payload.Success = true

	var err error
	payload.Domain, err = LoadDomain(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error = "" + err.Error()
	}
	payload.Mailboxes, err = GetMailboxes(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error = "" + err.Error()
	}
	payload.Aliases, err = GetAliases(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error = "" + err.Error()
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
