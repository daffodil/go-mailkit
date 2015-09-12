

package postfixadmin

import(
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


var Dbo gorm.DB

func SetupDb( engine string, db *sql.DB){

	// This is bummer cos I want to use db.Driver.Name or alike instead of a new function var
	var err error
	Dbo, err = gorm.Open(engine, db)

}


func SetupRoutes( router *mux.Router){

	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)

	//router.HandleFunc("/ajax/domain/{domain}/mailboxes", MailboxesAjaxHandler)

	//router.HandleFunc("/ajax/domain/{domain}/mailbox/{username}", MailboxAjaxHandler)
	//router.HandleFunc("/ajax/mailbox/{username}", MailboxAjaxHandler)

}
