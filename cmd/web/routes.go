package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (app *App) routes() fasthttp.RequestHandler {
	router := gin.Default()

	router.GET("/", app.Home)
	router.GET("/virtual-terminal", app.VirtualTerminal)
	router.POST("/payment-succeeded", app.PaymentSucceeded)
	return fasthttpadaptor.NewFastHTTPHandler(router)
}
