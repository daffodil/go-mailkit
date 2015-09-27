

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

	router.HandleFunc("/ajax/domains", AjaxHandlerDomains)
	router.HandleFunc("/ajax/domain/{domain}", AjaxHandlerDomain)
	router.HandleFunc("/ajax/domain/{domain}/all", AjaxHandlerDomainAll)
	router.HandleFunc("/ajax/domain/{domain}/vacations", AjaxHandlerVacations)
	router.HandleFunc("/ajax/domain/{domain}/mailboxes", AjaxHandlerMailboxes)
	router.HandleFunc("/ajax/domain/{domain}/virtual", AjaxHandlerDomainVirtual)

	router.HandleFunc("/ajax/domain/{domain}/mailbox/{username}", AjaxHandlerMailbox)

	router.HandleFunc("/ajax/mailbox/{email}", AjaxHandlerMailbox)
	router.HandleFunc("/ajax/mailbox/{email}/vacation", AjaxHandlerVacation)

	router.HandleFunc("/ajax/alias/{email}", AjaxHandlerAlias)
	router.HandleFunc("/ajax/domain/{domain}/aliases", AjaxHandlerAliases)





}
