

package mailconfig

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

