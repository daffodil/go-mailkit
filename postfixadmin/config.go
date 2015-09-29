

package postfixadmin

var conf *Config

type Config struct {

	Debug bool `yaml:"debug" json:"debug" `

	AuthSecret string `yaml:"auth_secret" json:"auth_secret" `

	Db DbConf

	VacationDomain string `yaml:"vacation_domain" json:"vacation_domain" `

	HTTPListen string `yaml:"http_listen" json:"http_listen"`
	IMAPAddress string `yaml:"imap_adddress" json:"imap_adddress"`
	SMTPLogin string `yaml:"smtp_login" json:"smtp_login"`

	//Tls *tls.Config
}


type DbConf struct {
	Engine string ` yaml:"engine" json:"engine" `
	Datasource string `yaml:"datasource" json:"connect"`
	Debug bool `yaml:"debug" json:"debug"`
	//Db *sql.DB

	TableNames map[string]string  `yaml:"table_names" json:"table_names"`
}





