package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"html/template"
	"log"
	"myapp/internal/driver"
	"myapp/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const version = "1.0.0"
const cssVersion = "1"

var session *scs.SessionManager

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB            models.DBModel
	Session       *scs.SessionManager
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf("Server run on Port: %v in mode: %s\n", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {

	gob.Register(map[string]interface{}{})
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment { development | production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.StringVar(&cfg.stripe.key, "stripe_key", "pk_test_51NJZOaAyXdxpP49B1kzBgfW9EK2YaGNtLKp2Ru4TRfugIzlIdtiGznzUOIY07w5IMFIiD1WGzV36HMBSGVLJJgCk00javhsPEb", "Stripe Key") //os.Getenv("STRIPE_KEY")

	flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch5:Schlucht6@tcp(db8.hostpark.net)/schmidschluch5?parseTime=true", "DB connect String")

	flag.Parse()

	cfg.stripe.key = "pk_test_51NJZOaAyXdxpP49B1kzBgfW9EK2YaGNtLKp2Ru4TRfugIzlIdtiGznzUOIY07w5IMFIiD1WGzV36HMBSGVLJJgCk00javhsPEb" //os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = "sk_test_51NJZOaAyXdxpP49B9HxMQsBwPMzqIKBRpv3cH4JFl1xEKRzfqBY8W3xKYEaAqkUmtn3RQUrCgESQKfZDa1QA3YOs007GqnxEu9"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
		return
	}
	defer conn.Close()

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		DB:            models.DBModel{DB: conn},
		Session:       session,
	}
	//app.infoLog.Println(cfg.stripe.key)
	//app.infoLog.Println(os.Getenv("STRIPE_KEY"))

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
