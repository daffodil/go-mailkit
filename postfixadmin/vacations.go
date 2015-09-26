

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"errors"

	"github.com/gorilla/mux"
)



type VacationsPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Vacations []*Vacation `json:"vacations"`
	Error string `json:"error"`
}



func GetVacations(domain string) ([]*Vacation, error) {
	var rows []*Vacation
	var err error

	if DomainExists(domain) == false {
		return rows, errors.New("Domain `" + domain + "` does not exist")
	}

	Dbo.Where("domain=?", domain).Find(&rows)
	return rows, err
}

// Handles /ajax/domain/<domain>/vacations
func VacationsAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("VacationsAjaxHandler")
	vars := mux.Vars(req)

	payload := VacationsPayload{}
	payload.Success = true
	payload.Vacations = make([]*Vacation, 0)

	var err error
	payload.Vacations, err = GetVacations( vars["domain"] )
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))

}


