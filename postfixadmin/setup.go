

package postfixadmin

import(
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)


var Dbx *sqlx.DB

func SetupDb( engine string, db *sql.DB){

	// This is bummer cos I want to use db.Driver.Name or alike instead of a new function var
	Dbx = sqlx.NewDb(db, engine)

}


func SetupRoutes( router *mux.Router){

	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/mailboxes", MailboxesAjaxHandler)

	router.HandleFunc("/ajax/domain/{domain}/mailbox/{username}", MailboxAjaxHandler)
	router.HandleFunc("/ajax/mailbox/{username}", MailboxAjaxHandler)

}
