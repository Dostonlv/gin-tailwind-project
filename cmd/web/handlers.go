package main

import (
	"github.com/gin-gonic/gin"
)

func (app *App) Home(c *gin.Context) {
	if err := app.renderTemplate(c, "home", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}

}

func (app *App) VirtualTerminal(c *gin.Context) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.Config.Stripe.Key
	if err := app.renderTemplate(c, "terminal", &templateData{
		StringMap: stringMap,
	}, "stripe-js"); err != nil {
		app.ErrorLog.Println(err)
	}
}
