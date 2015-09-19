

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"errors"

	"github.com/gorilla/mux"
)



type AliasesPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Aliases []Alias `json:"aliases"`
	Error string `json:"error"`
}


func GetAliases(domain string) ([]Alias, error) {

	var rows []Alias
	if DomainExists(domain) == false {
		return rows, errors.New("Domain '" + domain + "` does not exist")
	}
	var err error
	Dbo.Where("domain=?", domain).Order("address").Find(&rows)
	return rows, err
}

// Handles /ajax/domain/<domain>/aliases
func AliasesAjaxHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("AliasesAjaxHandler")
	vars := mux.Vars(req)

	payload := AliasesPayload{}
	payload.Success = true
	payload.Aliases = make([]Alias, 0)

	var err error
	payload.Aliases, err = GetAliases( vars["domain"] )
	if err != nil{
		fmt.Println(err)
		payload.Error = "Error: " + err.Error()
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
