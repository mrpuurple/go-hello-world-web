package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

<<<<<<< HEAD
	"github.com/mrpuurple/go-hello-world-web/pkg/config"
	"github.com/mrpuurple/go-hello-world-web/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the templates package
=======
	"github.com/mrpuurple/mygoapp/pkg/config"
	"github.com/mrpuurple/mygoapp/pkg/models"
)

// var functions = template.FuncMap{

// }

var app *config.AppConfig

// NewTemplates sets the config for the template package
>>>>>>> 888a37c (refactor)
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

<<<<<<< HEAD
// RenderTemplate renders a template to html
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the AppConfig
=======
// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template 

	if app.UseCache {
		// get the template cache from app config
>>>>>>> 888a37c (refactor)
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
<<<<<<< HEAD
		log.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from the ./templates folder
=======
		log.Println(err)
	}
}

// more complex way
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the file named *.page.tmpl from ./templates
>>>>>>> 888a37c (refactor)
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
<<<<<<< HEAD
		ts, err := template.New(name).ParseFiles(page)
=======
		ts, err := template.New(name).ParseFiles(page )
>>>>>>> 888a37c (refactor)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

<<<<<<< HEAD
		myCache[name] = ts
	}

	return myCache, nil
}
=======
		myCache[name] = ts  
	}

	return myCache, nil

}

// // simpler way
// // Create a map for template cache
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache (map)
// 	tc[t] = tmpl

// 	return nil
// }
>>>>>>> 888a37c (refactor)
