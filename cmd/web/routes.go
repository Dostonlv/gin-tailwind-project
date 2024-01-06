package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) routes() http.Handler {
	route := gin.Default()

	route.GET("/", app.Home)

	return route

}
