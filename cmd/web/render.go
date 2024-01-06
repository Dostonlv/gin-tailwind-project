package main

import (
	"embed"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *App) addDefaultData(td *templateData, c *gin.Context) *templateData {
	td.API = app.Config.API
	return td
}

func (app *App) renderTemplate(c *gin.Context, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.TemplateCache[templateToRender]

	if app.Config.ENV == "production" && templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, c)
	err = t.Execute(c.Writer, td)
	if err != nil {
		app.ErrorLog.Println(err)
		return err
	}

	return nil
}

func (app *App) parseTemplate(partials []string, page string, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}
	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}

	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}
	app.TemplateCache[templateToRender] = t
	return t, nil
}
