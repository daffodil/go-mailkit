

package mailadmin

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


func NewMailboxesPayload() MailboxesPayload {
	t := MailboxesPayload{}
	t.Success = true
	t.Mailboxes = make([]Mailbox, 0)
	return t
}


func GetMailboxes() ([]Domain, error) {
	var rows []Mailbox
	var err error
	err = Db.Select(&rows, "SELECT username, password, name, maildir, quota, local_part, domain, created, modified, active FROM domain order by username asc ")
	return rows, err
}

// Handles /ajax/domain/<domain>/mailboxes
func MailboxesAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	//_ := mux.Vars(req)
	// TODO auth

	payload := NewDomainsPayload()

	var err error
	payload.Domains, err = GetMailboxes()
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
