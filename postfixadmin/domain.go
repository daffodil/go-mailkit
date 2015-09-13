

package postfixadmin

import(

	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Domain struct {
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

func(me Domain) TableName() string {
	return TableNames["domain"]
}

// The ajax data returned for a domain
type DomainPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Domain Domain `json:"domain"`
	Error string `json:"error"`
}


// Create the return struct for a domain
func CreateDomainPayload() DomainPayload {
	payload := DomainPayload{}
	payload.Success = true
	payload.Domain = Domain{}
	return payload
}


func GetDomain(domain_name string) (Domain, error) {

	var dom Domain
	var err error
	Dbo.Where("domain = ? ", domain_name).Find(&dom)
	//err := row.Scan(&dom.Domain, &dom.Description, &dom.Aliases, &dom.Mailboxes, &dom.MaxQuota, &dom.Quota, &dom.Transport, &dom.BackupMx, &dom.Created, &dom.Modified, &dom.Active)
	return dom, err
}


// Returns json at  /ajax/domain/{domain}
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
