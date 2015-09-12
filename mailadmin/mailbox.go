

package mailadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Mailbox struct {
	Usernname string 	`json:"domain"`
	Password string 	`json:"password"`
	Name string 		`json:"name"`
	Maildir int 		`json:"maildir"`
	Quota int 			`json:"quota"`
	LocalPart int 		`json:"local_part"`
	Domain string		`json:"domain"`
	Created string		`json:"created"`
	Modified string		`json:"modified"`
	Active int 			`json:"active"`
}

type MailboxPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Mailbox Mailbox `json:"mailbox"`
	Error string `json:"error"`
}



func CreateMailboxPayload() MailboxPayload {
	payload := MailboxPayload{}
	payload.Success = true
	payload.Mailbox = Mailbox{}
	return payload
}


func GetMailbox(username string) (Mailbox, error) {
	var dom Mailbox
	row := Db.QueryRow("SELECT domain, description, aliases, mailboxes, maxquota, quota, transport, backupmx, created, modified, active FROM domain where domain = ? ", username)
	err := row.Scan(&dom.Usernname, &dom.Password, &dom.Name, &dom.Maildir, &dom.Quota, &dom.LocalPart, &dom.Domain, &dom.Created, &dom.Modified, &dom.Active)
	return dom, err
}


// Handles /ajax/domain/<example.com>
func MailboxAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	payload := CreateMailboxPayload()

	var err error
	payload.Mailbox, err = GetDomain(vars["username"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
