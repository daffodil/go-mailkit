

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Domain struct {
	//DomainID int `db:"domain_id" json:"domain_id"`
	Domain string 		`db:"domain" json:"domain"`
	Description string 	`db:"description" json:"description"`
	Aliases int 		`db:"aliases" json:"aliases"`
	Mailboxes int 		`db:"mailboxes" json:"mailboxes"`
	MaxQuota int 		`db:"maxquota" json:"maxquota"`
	Quota int 			`db:"quota" json:"quota"`
	Transport string	`db:"transport" json:"transport"`
	BackupMx int 		`db:"backupmx" json:"backupmx"`
	Created string		`db:"created" json:"created"`
	Modified string		`db:"modified" json:"modified"`
	Active int 			`db:"active" json:"active"`
}

type DomainPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Error string `json:"error"`
}



func CreateDomainPayload() DomainPayload {
	payload := DomainPayload{}
	payload.Success = true
	payload.Domain = Domain{}
	return payload
}



func GetDomain(domain_name string) (Domain, error) {
	var dom Domain
	row := Dbx.QueryRow("SELECT domain, description, aliases, mailboxes, maxquota, quota, transport, backupmx, created, modified, active FROM domain where domain = ? ", domain_name)
	err := row.Scan(&dom.Domain, &dom.Description, &dom.Aliases, &dom.Mailboxes, &dom.MaxQuota, &dom.Quota, &dom.Transport, &dom.BackupMx, &dom.Created, &dom.Modified, &dom.Active)
	return dom, err
}


// Handles /ajax/domain/<example.com>
func DomainAjaxHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("DomainAjaxHandler")
	vars := mux.Vars(req)

	payload := CreateDomainPayload()

	var err error
	payload.Domain, err = GetDomain(vars["domain"])
	if err != nil{
		fmt.Println(err)
		payload.Error = "DB Error: " + err.Error()
	}


	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
