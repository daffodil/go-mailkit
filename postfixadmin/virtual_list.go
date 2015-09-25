

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	//"errors"

	"github.com/cenkalti/log"
	"github.com/gorilla/mux"
)



//= Ajax struct for `domain` all
type DomainVirtualPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Mailboxes []Mailbox `json:"mailboxes"`
	Aliases []Alias `json:"aliases"`
	Error string `json:"error"`
}


//=  /ajax/domain/{domain}/all
func DomainVirtualAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	log.Info("DomainVirtualAjaxHandler")

	vars := mux.Vars(req)
	domain := vars["domain"]

	payload := DomainAllPayload{}
	payload.Success = true

	var err error
	payload.Domain, err = GetDomain(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error =  err.Error()
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
