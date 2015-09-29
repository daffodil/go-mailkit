

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
	Active bool 		`json:"active"`
}

func(me Mailbox) TableName() string {
	return conf.Db.TableNames["mailbox"]
}





func LoadMailbox(username string) (Mailbox, error) {
	var mailbox Mailbox
	var err error
	Dbo.Where("username = ? ", username).First(&mailbox)

	return mailbox, err
}

func MailboxExists(address string ) bool {
	var count int
	Dbo.Model(Mailbox{}).Where("username = ?", address).Count(&count)
	if count == 0 {
		return false
	}
	return true
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


// /ajax/domain/<example.com>/mailbox/<email>
func AjaxHandlerMailbox(resp http.ResponseWriter, req *http.Request) {


	if AjaxAuth(resp, req) == false {
		return
	}

	vars := mux.Vars(req)

	payload := CreateMailboxPayload()

	var err error
	payload.Mailbox, err = LoadMailbox(vars["username"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}

	switch req.Method {

	case "POST":

		f := req.Form
		//fmt.Println(f)
		payload.Vacation.Mailbox = email_addr.Address
		payload.Vacation.Domain = email_addr.Domain
		payload.Vacation.Active, err = strconv.ParseBool(f.Get("active"))
		payload.Vacation.Activefrom = f.Get("active_from")
		payload.Vacation.Activeuntil = f.Get("active_until")
		payload.Vacation.IntervalTime, err = strconv.ParseInt(f.Get("interval_time"), 10, 64)
		payload.Vacation.Subject = f.Get("subject")
		payload.Vacation.Body = f.Get("body")

		Dbo.Save(&payload.Vacation)
		fmt.Println("------------------ POSTED-------------")
		UpdateVacationAlias(payload.Vacation)


	case "GET":

		// just pass through vars

	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
