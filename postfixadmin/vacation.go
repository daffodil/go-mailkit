

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"strings"

	"github.com/gorilla/mux"
)


type Vacation struct {
	Email string 		` json:"email" `
	Subject string 		` json:"subject" `
	Body string 		` json:"body" `
	ActiveFrom string 	` json:"active_from" `
	ActiveTo string 	` json:"active_to" `
	Cache string 		` json:"cache" `
	Domain string 		` json:"domain" `
	IntervalTime string ` json:"interval_time" `
	Created string 		` json:"created" `
	Modified string 	` json:"modified" `
	Active bool		 	` json:"active" `
}



type VacationNotification struct {
	OnVacation string 	` json:"on_vacation" `
	Notified string 	` json:"notified" `
	NotifiedAt string 	` json:"notified_at" `
}


type VacationPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Vacation *Vacation `json:"vacation"`
	Error string `json:"error"`
}


func IsVacationAddress(address string) bool {

	user_domain :=  strings.Split(address, "@")
	if user_domain[1] == VacationDomain {
		return true
	}
	return false
}

func GetVacation(email string) (*Vacation, error) {
	row := new(Vacation)
	var err error
	Dbo.Where("email = ?", email).Find(&row)
	if row.Email == "" {
		// no record so return
		return nil, err
	}
	return row, err
}

// Handles /ajax/vacation/<email>
func VacationAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("VacationsAjaxHandler")

	payload := VacationPayload{}
	payload.Success = true //extjs fu

	vars := mux.Vars(req)

	email_addr, err_email := ParseAddress(vars["email"])
	if err_email != nil {
		payload.Error = err_email.Error()
	} else {

		// check mail exists



		var err error
		payload.Vacation, err = GetVacation(email_addr.Address)
		if err != nil {
			fmt.Println(err)
			payload.Error = "DB Error: "+err.Error()
		}
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}


