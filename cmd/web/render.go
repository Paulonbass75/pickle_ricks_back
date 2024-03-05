package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float64
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated bool
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.config.api
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	tempToRender := fmt.Sprintf("templates/%s.page.tmpl", page)

	_, templateInMap := app.templateCache[tempToRender]

	if app.config.env == "production" && templateInMap {
		t = app.templateCache[tempToRender]
	} else {
		t, err = app.parseTemplate(partials, page, tempToRender)
		if err != nil {
			app.errorlog.Println(err)
			return err

		}
	}

	if td == nil {
        td = &templateData{}
    }
	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err!= nil {
        app.errorlog.Println(err)
        return err
    }

	return nil

}

func (app *application) parseTemplate(partials []string, page string, tempToRender string) (*template.Template, error) {
	var t *template.Template
	var err error
	//build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.tmpl", x)
		}
	}
	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", strings.Join(partials, "" ),tempToRender)
	} else {
	t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", tempToRender)
	}
	if err!= nil {
        app.errorlog.Println(err)
        return nil, err
    }
	app.templateCache[tempToRender] = t
	return t, nil
}
