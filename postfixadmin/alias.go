

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)


type Alias struct {
	Address string 		`db:"address" json:"address"`
	Goto string 		`db:"goto" json:"goto"`
	Domain string 		`db:"domain" json:"domain"`
	Created string		`db:"created" json:"created"`
	Modified string		`db:"modified" json:"modified"`
	Active int			`db:"active" json:"active"`
}

func(me Alias) TableName() string {
	return TableNames["alias"]
}

type AliasPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Alias []Alias `json:"alias"`
	Error string `json:"error"`
}



func CreateAliasPayload() AliasPayload {
	payload := AliasPayload{}
	payload.Success = true
	//payload.Alias = make(Alias, 0)
	return payload
}



func GetAlias(email string) ([]Alias, error) {

	var alias []Alias
	var err error
	Dbo.Where("address = ? ", email).Find(&alias)
	return alias, err
}

// Handles /ajax/alias/<email>
func AliasAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("AliasAjaxHandler")
	vars := mux.Vars(req)

	payload := CreateAliasPayload()

	var err error
	payload.Alias, err = GetAlias(vars["email"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
