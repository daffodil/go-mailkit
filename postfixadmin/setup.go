

package postfixadmin

import(
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"



)

var TableNames  map[string]string
var VacationDomain string


var Dbo gorm.DB

func Initialize( db_engine string, db *sql.DB, table_names map[string]string, sql_log bool, vacation_domain string){

	// This is bummer cos I want to use db.Driver.Name or alike instead of a new function var
	var err error
	Dbo, err = gorm.Open(db_engine, db)
	if err != nil {

	}
	Dbo.SingularTable(true)
	Dbo.LogMode(sql_log)

	TableNames = table_names
	VacationDomain = vacation_domain
}


func SetupRoutes(router *mux.Router){

	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/all", DomainAllAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/vacations", VacationsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/mailboxes", MailboxesAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/virtual", DomainVirtualAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/mailbox/{username}", MailboxAjaxHandler)

	router.HandleFunc("/ajax/mailbox/{email}", MailboxAjaxHandler)
	router.HandleFunc("/ajax/mailbox/{email}/vacation", VacationAjaxHandler)

	router.HandleFunc("/ajax/alias/{email}", AliasAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/aliases", AliasesAjaxHandler)





}
