

package mailadmin

import(

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)


var Db *sqlx.DB

func Setup( db *sqlx.DB, router *mux.Router){
	Db = db
	router.HandleFunc("/ajax/domains", DomainsAjaxHandler)
	router.HandleFunc("/ajax/domain/{domain}", DomainAjaxHandler)



}
