package config

import (
	"html/template"
	"log"
)

const Version = "1.0.0"
const CSSVersion = "1"

type Config struct {
	Port string
	ENV  string
	API  string
	DB   struct {
		DSN string
	}
	Stripe struct {
		Secret string
		Key    string
	}
}

type Application struct {
	Config        Config
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	TemplateCache map[string]*template.Template
	Version       string
}
