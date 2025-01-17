package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"go-breeders/adapters"
	"go-breeders/configuration"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useCache       bool
	dataSourceName string
}

func main() {
	app := new(application)

	app.templateMap = make(map[string]*template.Template)

	flag.StringVar(&app.config.dataSourceName, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.Parse()

	// get database
	db, err := initMySQLDB(app.config.dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	// jsonBackend := &adapters.JSONBackend{}
	// jsonAdapter := &adapters.RemoteService{Remote: jsonBackend}

	xmlBackend := &adapters.XMLBackend{}
	xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	app.App = configuration.New(db, xmlAdapter)

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("starting web application on port", port)

	log.Fatal(server.ListenAndServe())
}
