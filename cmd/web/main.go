package main

import (
	"flag"
	"github.com/Dostonlv/gin-tailwind-project/config"
	"github.com/valyala/fasthttp"
	"html/template"

	"log"

	"os"

	"time"
)

type App config.Application

func main() {
	var cfg config.Config

	flag.StringVar(&cfg.Port, "port", "4000", "Server port to listen on")
	flag.StringVar(&cfg.ENV, "env", "development", "Application enviornment {development|production}")
	flag.StringVar(&cfg.DB.DSN, "dsn", "itachi:secret@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.API, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := App{
		Config:        cfg,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		TemplateCache: tc,
		Version:       config.Version,
	}

	err := app.serve()

	if err != nil {
		app.ErrorLog.Fatal(err)
	}

}

func (app *App) serve() error {
	srv := &fasthttp.Server{
		Handler:      app.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	if app.Config.ENV == "development" {
		app.InfoLog.Printf("Starting server on port %s\n", app.Config.Port)
		return srv.ListenAndServe(":" + app.Config.Port)
	}

	app.InfoLog.Printf("Starting HTTP server in %s mode on port %s\n", app.Config.ENV, app.Config.Port)

	return nil

}
