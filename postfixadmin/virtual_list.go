

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	//"errors"
	"strings"

	"github.com/cenkalti/log"
	"github.com/gorilla/mux"
)

type MailboxVirtual struct {
	Mailbox
	Vacation *Vacation		`json:"vacation"`
	ForwardOnly bool 		`json:"forward_only"`
	ForwardFrom []string	`json:"forward_from"`
	ForwardTo []string		`json:"forward_to"`
}


//= Ajax struct for `domain` all
type DomainVirtualPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Mailboxes []*MailboxVirtual `json:"mailboxes"`
	Aliases []Alias `json:"aliases"`
	Error string `json:"error"`
}


func GetMailboxesVirtual(domain string) ([]*MailboxVirtual, error) {
	var rows []*MailboxVirtual
	var err error
	Dbo.Where("domain=?", domain).Order("username asc").Find(&rows)
	return rows, err
}

//=  /ajax/domain/{domain}/all
func DomainVirtualAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	log.Info("DomainVirtualAjaxHandler")

	vars := mux.Vars(req)
	domain := vars["domain"]

	payload := DomainVirtualPayload{}
	payload.Success = true

	var err error
	payload.Domain, err = GetDomain(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error =  err.Error()
	}

	// Get mailboxes with virtual struct
	payload.Mailboxes, err = GetMailboxesVirtual(domain)
	if err != nil {
		log.Info(err.Error())
		payload.Error =  err.Error()
	}

	// make mailboxes map with user > MailboxVirt and init mbox
	mailboxes_map := make(map[string]*MailboxVirtual)
	for _, mb := range payload.Mailboxes {
		mb.ForwardFrom = make([]string, 0)
		mb.ForwardTo = make([]string, 0)
		mb.ForwardOnly = true
		mailboxes_map[mb.Username] = mb
	}
	//fmt.Println(mb_map)

	// Load the Vacations
	vacations, errv := GetVacations(domain)
	if errv != nil {
		log.Info(errv.Error())
		payload.Error =  err.Error()
	}
	for _, va := range vacations {
		mailboxes_map[va.Email].Vacation = va
	}

	// Get list of aliases and postfix
	// does curious things such as forwarding to same mailbox
	payload.Aliases, err = GetAliases(domain)
	if err != nil{
		log.Info(err.Error())
		payload.Error = "" + err.Error()
	}
	for _, alias := range payload.Aliases {
		mbox, ok := mailboxes_map[alias.Address]
		if ok {
			gotos := SplitEmail(alias.Goto)
			fmt.Println(mbox.Username, gotos)
			for _, gotto := range gotos {
				if gotto == mbox.Username {
					mbox.ForwardOnly = false
				} else {
					mbox.ForwardTo = append(mbox.ForwardTo, gotto)
				}

				mbox2, ok2 := mailboxes_map[gotto]
				if ok2 {
					//if mbox2.Username != alias.Address {
						mbox2.ForwardFrom = append(mbox.ForwardFrom, gotto)
					//}
				}
			}
		}
	}




	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}


func SplitEmail(email_str string) []string {
	parts := strings.Split(email_str, ",")
	return parts
}
