
package mailconfig

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type DbConf struct {
	Engine string ` yaml:"engine" json:"engine" `
	Connect string `yaml:"connect" json:"connect"`
	Debug bool `yaml:"debug" json:"debug"`
	Db *sql.DB
}


