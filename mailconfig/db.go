
package mailconfig

import (
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)


type DbConf struct {
	Engine string ` yaml:"engine" json:"engine" `
	Connect string `yaml:"connect" json:"connect"`
	SqlDebug bool `yaml:"debug" json:"debug"`
	//Db *sql.DB

	TableNames map[string]string  `yaml:"table_names" json:"table_names"`
}


