package main

import (
	"github.com/Dostonlv/gin-tailwind-project/internal/models"
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

func (app *App) PaymentSucceeded(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
	cardHolder := c.Request.Form.Get("cardholder_name")
	email := c.Request.Form.Get("cardholder_email")
	paymentIntent := c.Request.Form.Get("payment_intent")
	paymentMethod := c.Request.Form.Get("payment_method")
	paymentAmount := c.Request.Form.Get("payment_amount")
	paymentCurrency := c.Request.Form.Get("payment_currency")
	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["cardholder_email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency
	if err := app.renderTemplate(c, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.ErrorLog.Println(err)
	}
}

func (app *App) ChargeOnce(c *gin.Context) {
	stringMap := make(map[string]string)
	widget := models.Widget{
		ID:             1,
		Name:           "Custom Widget",
		Description:    "Very nice widget",
		InventoryLevel: 10,
		Price:          1000,
	}

	data := make(map[string]any)
	data["widget"] = widget

	stringMap["publishable_key"] = app.Config.Stripe.Key
	if err := app.renderTemplate(c, "buy-once", &templateData{
		StringMap: stringMap,
		Data:      data,
	}, "stripe-js"); err != nil {

		app.ErrorLog.Println(err)
	}
}
