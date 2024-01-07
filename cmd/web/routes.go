package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (app *App) routes() fasthttp.RequestHandler {
	router := gin.Default()

	router.GET("/", app.Home)
	return fasthttpadaptor.NewFastHTTPHandler(router)
}
