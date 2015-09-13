

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)


type MailboxesPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Mailboxes []Mailbox `json:"mailboxes"`
	Error string `json:"error"`
}


func CreateMailboxesPayload() MailboxesPayload {
	t := MailboxesPayload{}
	t.Success = true
	t.Mailboxes = make([]Mailbox, 0)
	return t
}


func GetMailboxes(domain string) ([]Mailbox, error) {
	var rows []Mailbox
	var err error
	Dbo.Where("domain=?", domain).Find(&rows)
	return rows, err
}

// Handles /ajax/domain/<domain>/mailboxes
func MailboxesAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("MailboxesAjaxHandler")
	vars := mux.Vars(req)

	payload := CreateMailboxesPayload()

	var err error
	payload.Mailboxes, err = GetMailboxes( vars["domain"] )
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}