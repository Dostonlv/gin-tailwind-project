package main

import (
	"encoding/json"
	"strconv"

	"github.com/Dostonlv/gin-tailwind-project/internal/cards"
	"github.com/Dostonlv/gin-tailwind-project/internal/models"
	"github.com/gin-gonic/gin"
)

func (app *App) GetPaymentIntent(c *gin.Context) {
	var payload models.StripePayload

	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	cars := cards.Card{
		Secret:   app.Config.Stripe.Secret,
		Key:      app.Config.Stripe.Key,
		Currency: payload.Currency,
	}

	okay := true
	pi, msg, err := cars.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}

	if okay {
		out, err := json.MarshalIndent(pi, "", "    ")
		if err != nil {
			app.ErrorLog.Println(err)
			return
		}
		c.Header("Content-Type", "application/json")
		c.Writer.Write(out)

	} else {
		j := models.JsonResponse{
			OK:      false,
			Message: msg,
			Content: "",
		}
		out, err := json.MarshalIndent(j, "", "   ")
		if err != nil {
			app.ErrorLog.Println(err)
			return
		}
		c.Header("Content-Type", "application/json")
		c.Writer.Write(out)
	}

}

func (app *App) GetWidgetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	widget, err := app.DB.GetWidget(id)
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	out, err := json.MarshalIndent(widget, "", "    ")
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Writer.Write(out)
}
