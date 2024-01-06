package main

import (
	"github.com/gin-gonic/gin"
)

func (app *App) Home(c *gin.Context) {
	if err := app.renderTemplate(c, "home", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}

}
