

package postfixadmin

import(
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var TableNames  map[string]string



var Dbo gorm.DB

func SetupDb( engine string, db *sql.DB, table_names map[string]string){

	// This is bummer cos I want to use db.Driver.Name or alike instead of a new function var
	var err error
	Dbo, err = gorm.Open(engine, db)
	if err != nil {

	}
	Dbo.SingularTable(true)
	Dbo.LogMode(true)

	TableNames = table_names

	DomainsMap = make(map[string]Domain)

	LoadDomainsMap()
}


func SetupRoutes( router *mux.Router){

	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/mailboxes", MailboxesAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/mailbox/{username}", MailboxAjaxHandler)

	router.HandleFunc("/ajax/mailbox/{email}", MailboxAjaxHandler)
	router.HandleFunc("/ajax/mailbox/{email}/vacation", VacationAjaxHandler)

	router.HandleFunc("/ajax/alias/{email}", AliasAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}/aliases", AliasesAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/vacations", VacationsAjaxHandler)


}
