package main

import (
	"flag"
	"github.com/Dostonlv/gin-tailwind-project/config"
	"github.com/Dostonlv/gin-tailwind-project/internal/driver"
	"log"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

type App config.Application

func main() {
	var cfg config.Config
	flag.StringVar(&cfg.Port, "port", "4001", "Server port to listen on")
	flag.StringVar(&cfg.DB.DSN, "dsn", "itachi:secret@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.ENV, "env", "development", "Application enviornment {development|production|maintenance}")
	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conn, err := driver.OpenDB(cfg.DB.DSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &App{
		Config:   cfg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Version:  config.Version,
	}
	err = app.serve()
	if err != nil {
		app.ErrorLog.Printf("Error starting server: %v\n", err)
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
