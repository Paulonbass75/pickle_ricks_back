package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"pickle_ricks_back/internal/driver"
	"pickle_ricks_back/models"
	"time"

	"github.com/joho/godotenv"
)

const version = "1.0.0"
// const cssVersion = "1"

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
	errorlog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB models.DBModel
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

	app.infoLog.Printf(fmt.Sprintf("Starting HTTP server in %s mode on port %d\n", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development | production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "paul:password1@tcp(localhost:3306)/pickle_test?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.api, "api", "http://localhost:8081", "URL to API")

	flag.Parse()

	if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
}

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err!= nil {
        log.Fatal(err)
    }
	defer conn.Close()

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorlog:      errorLog,
		templateCache: tc,
		version:       version,
		DB: models.DBModel{DB: conn},
	}
	err = app.serve()
	if err != nil {
		app.errorlog.Println(err)
		log.Fatal(err)
	}
}
