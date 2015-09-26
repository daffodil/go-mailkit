

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
	LocalPart string	`json:"local_part"`
	Domain string		`json:"domain"`
	Created string		`json:"created"`
	Modified string		`json:"modified"`
	Active bool 			`json:"active"`
}

func(me Mailbox) TableName() string {
	return TableNames["mailbox"]
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
	var mailbox Mailbox
	var err error
	Dbo.Where("username = ? ", username).First(&mailbox)

	return mailbox, err
}


// Handles /ajax/domain/<example.com>/mailbox/<email>
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
