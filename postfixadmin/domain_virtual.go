

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
	aliases, err_al := GetAliases(domain)
	if err_al != nil{
		log.Info(err_al.Error())
		payload.Error = "" + err_al.Error()
	}
	//aliases_map := make(map[string][]string)

	for _, alias := range aliases {


		mbox, ok := mailboxes_map[alias.Address]
		if ok {
			// Alias is also a mailbox - postfix style
			gotos := SplitEmail(alias.Goto)

			for _, gotto := range gotos {

				// postfix forwards alias to mailbox
				if gotto == mbox.Username {
					mbox.ForwardOnly = false

				// User on vactaion
				} else if IsVacationAddress(gotto){
					// to nothing

				} else {
					mbox.ForwardTo = append(mbox.ForwardTo, gotto)

				}

				if gotto != mbox.Username {
					targetBox, ok2 := mailboxes_map[gotto]
					if ok2 {
						//fmt.Println("----", alias.Address, gotto, targetBox.Username, aliases_map)

						targetBox.ForwardFrom = append(targetBox.ForwardFrom, mbox.Username)

					}
				}

			}



		} else {
			fmt.Println("========", alias)
			gotos2 := SplitEmail(alias.Goto)
			for _, got2 := range gotos2 {
				targetBox2, ok3 := mailboxes_map[got2]
				if ok3 {
					targetBox2.ForwardFrom = append(targetBox2.ForwardFrom, alias.Address)
				}
			}
			payload.Aliases = append( payload.Aliases, alias)
		}
	}




	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}


func SplitEmail(email_str string) []string {
	parts := strings.Split(email_str, ",")
	return parts
}
