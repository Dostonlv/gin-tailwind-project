package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (app *App) routes() fasthttp.RequestHandler {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"}
	config.AllowCredentials = false
	config.MaxAge = 300

	r.Use(cors.New(config))

	r.POST("/api/payment-intent", app.GetPaymentIntent)
	r.GET("/api/widget/:id", app.GetWidgetByID)
	return fasthttpadaptor.NewFastHTTPHandler(r)
}
