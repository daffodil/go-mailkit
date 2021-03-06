

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"
	"strings"

	"github.com/gorilla/mux"
)

// Represents the alias table
type Alias struct {
	Address string 		`json:"address" gorm:"primary_key"`
	Goto string 		`json:"goto"`
	Domain string 		`json:"domain"`
	Created string		`json:"created"`
	Modified string		`json:"modified"`
	Active int			`json:"active"`
}

func (me *Alias) TableName() string {
	return conf.Db.TableNames["alias"]
}

// Save instance to Db
func (me *Alias) Save(){
	Dbo.Save(&me)
}

// Add address to forwarding
func (me *Alias) AddGoto(addr string) {
	parts := strings.Split(me.Goto, ",")
	found := false
	for _, p := range parts {
		if p == addr {
			found = true
		}
	}
	if found == true {
		//fmt.Println("DOun vac alias")
		return
	}
	parts = append(parts, addr)
	me.Goto = strings.Join(parts, ",")
}


// Remove address from forwarding
func(me *Alias) RemoveGoto(addr string) {

	addresses := make([]string, 0)
	gotos := strings.Split(me.Goto, ",")

	for _, p := range gotos {
		if p != addr {
			addresses = append(addresses, p)
		}
	}
	me.Goto = strings.Join(addresses, ",")
}


func GetAlias(email string) (Alias, error) {

	var alias Alias
	var err error
	Dbo.Where("address = ? ", email).Find(&alias)
	return alias, err
}




type AliasPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Alias Alias `json:"alias"`
	Error string `json:"error"`
}

func CreateAliasPayload() AliasPayload {
	payload := AliasPayload{}
	payload.Success = true
	//payload.Alias = make(Alias, 0)
	return payload
}




// /ajax/alias/<email>
func AjaxHandlerAlias(resp http.ResponseWriter, req *http.Request) {
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
