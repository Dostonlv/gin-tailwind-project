package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

func (app *App) routes() fasthttp.RequestHandler {
	router := gin.Default()

	router.GET("/", app.Home)
	router.GET("/virtual-terminal", app.VirtualTerminal)
	router.POST("/payment-succeeded", app.PaymentSucceeded)
	router.GET("/charge-once", app.ChargeOnce)
	router.StaticFS("/static", http.Dir("./static"))

	return fasthttpadaptor.NewFastHTTPHandler(router)
}
