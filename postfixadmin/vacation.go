

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"errors"
	"strconv"

	"github.com/gorilla/mux"
)


type Vacation struct {
	Email string 		` json:"email" gorm:"primary_key" `
	Subject string 		` json:"subject" `
	Body string 		` json:"body" `
	Activefrom string 	` json:"active_from" `
	Activeuntil string 	` json:"active_until" `
	Cache string 		` json:"cache" `
	Domain string 		` json:"domain" `
	IntervalTime int64 ` json:"interval_time" `
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
		if !MailboxExists(email_addr.Address) {
			payload.Error = errors.New("Mailbox `" + email_addr.Address + "` does not exist").Error()

		} else {

			var err error
			payload.Vacation, err = GetVacation(email_addr.Address)
			if err != nil {
				fmt.Println(err)
				payload.Error = "DB Error: "+err.Error()
			}

			switch req.Method {

			case "POST":
				req.ParseForm()
				f := req.Form
				//fmt.Println(f)
				payload.Vacation.Email = email_addr.Address
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


				if payload.Vacation.Email == "" {
					// probably record not exist
					payload.Vacation.Email = email_addr.Address
				}
			}
		}
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}


func UpdateVacationAlias(vac *Vacation) {

	alias, err := GetAlias(vac.Email)
	fmt.Println("UpdateVacationAlias", alias, err)
	if err != nil {
		// do something
		return
	}
	em, errp := ParseAddress(vac.Email)
	if errp != nil {
		return
	}
	if vac.Active {
		alias.AddGoto(em.VacationAddress)
	}else{
		alias.RemoveGoto(em.VacationAddress)
	}
	alias.Save()
}
