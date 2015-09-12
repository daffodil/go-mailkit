
package main

import (
	"fmt"
	//"log"
	"flag"
	"net/http"
	"io/ioutil"
	"os"
	//"crypto/tls"

	"gopkg.in/yaml.v2"
	"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"


	"github.com/daffodil/go-mailkit/postfixadmin"

)


type Config struct {

	Debug bool `yaml:"debug" json:"debug" `

	AuthSecret string `yaml:"auth_secret" json:"auth_secret" `

	DBEngine string `yaml:"db_engine" json:"db_engine"`
	DBConnect string `yaml:"db_connect" json:"db_connect"`

	HTTPListen string `yaml:"http_listen" json:"http_listen"`
	IMAPAddress string `toml:"imap_adddress" json:"imap_adddress"`
	SMTPLogin string `toml:"smtp_login" json:"smtp_login"`

	//Tls *tls.Config
}


func main(){

	config_file := flag.String("config", "config.yaml", "Config file")
	flag.Parse()

	// Create and load config.yaml
	config := new(Config)
	contents, e := ioutil.ReadFile(*config_file)
	if e != nil {
		fmt.Printf("Config File Error: %v\n", e)
		fmt.Printf("create one with -w \n")
		os.Exit(1)
	}
	if err := yaml.Unmarshal(contents, &config); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Setup TLS config
	//config.Tls = new(tls.Config)
	//config.Tls.ServerName = config.IMAPAddress
	//config.Tls.InsecureSkipVerify = true

	fmt.Printf("Config: %v\n", config)

	// Create Database connection
	var Db *sql.DB
	var err_db error
	Db, err_db = sql.Open(config.DBEngine, config.DBConnect)
	if err_db != nil {
		fmt.Printf("Db Login Failed: ", err_db,"=", config.DBEngine)
		os.Exit(1)
	}
	err_ping := Db.Ping()
	if err_ping != nil {
		fmt.Printf("Db Ping Failed: ", err_ping,"=", config.DBEngine)
		os.Exit(1)
	}
	defer Db.Close()

	// Setup router and config mods
	router := mux.NewRouter()
	postfixadmin.SetupDb(config.DBEngine, Db)
	postfixadmin.SetupRoutes(router)
	//mailbox.Configure(config, r)


	fmt.Println("Serving on " + config.HTTPListen)
	http.Handle("/", router)
	http.ListenAndServe( config.HTTPListen, nil)

}

