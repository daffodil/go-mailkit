
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

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	//"github.com/daffodil/go-mailkit/mailconfig"
	"github.com/daffodil/go-mailkit/postfixadmin"

)




func main(){

	config_file := flag.String("config", "config.yaml", "Config file")
	flag.Parse()

	// Create and load config.yaml
	config := new(postfixadmin.Config)
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

	fmt.Printf("Config: %v\n", config.Db.TableNames)

	// Create Database connection
	var Db *sql.DB
	var err_db error
	Db, err_db = sql.Open(config.Db.Engine, config.Db.Connect)
	if err_db != nil {
		fmt.Printf("Db Login Failed: ", err_db,"=", config.Db.Engine, config.Db.Connect)
		os.Exit(1)
	}
	err_ping := Db.Ping()
	if err_ping != nil {
		fmt.Printf("Db Ping Failed: ", err_ping,"=", config.Db.Engine, config.Db.Connect)
		os.Exit(1)
	}
	defer Db.Close()
	postfixadmin.Initialize(config, Db)


	// Setup router and config mods
	router := mux.NewRouter()
	postfixadmin.SetupRoutes(router)



	fmt.Println("Serving on " + config.HTTPListen)
	http.Handle("/", router)
	http.ListenAndServe( config.HTTPListen, nil)

}

