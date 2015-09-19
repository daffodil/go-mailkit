

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"errors"

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
