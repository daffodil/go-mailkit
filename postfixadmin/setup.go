

package postfixadmin

import(
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

)


var Dbo gorm.DB

// Initializes the postfix admin module..
func Initialize( conff *Config, db *sql.DB){

	conf = conff

	// This is bummer cos I want to use db.Driver.Name or alike instead of a new function var
	var err error
	Dbo, err = gorm.Open(conf.Db.Engine, db)
	if err != nil {

	}
	Dbo.SingularTable(true)
	Dbo.LogMode(conf.Debug)

}


// Add routes for postfixadmin module. Idea is that
// if your not admin or alike, then u get a 404 or 500
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
