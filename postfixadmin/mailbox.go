

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Mailbox struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	Name string 		`json:"name"`
	Maildir string 		`json:"maildir"`
	Quota int 			`json:"quota"`
	Local_Part string	`json:"local_part"`
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
	row := Db.QueryRow("SELECT username, password, name, maildir, quota, localpart, domain, created, modified, active FROM mailbox where username = ? ", username)
	err := row.Scan(&dom.Username, &dom.Password, &dom.Name, &dom.Maildir, &dom.Quota, &dom.Local_Part, &dom.Domain, &dom.Created, &dom.Modified, &dom.Active)
	return dom, err
}


// Handles /ajax/domain/<example.com>
func MailboxAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("mailboxhandler")
	vars := mux.Vars(req)

	payload := CreateMailboxPayload()

	var err error
	payload.Mailbox, err = GetMailbox(vars["username"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
