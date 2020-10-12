package templates

import (
	"fmt"
	"html/template"
	"path/filepath"
	"tsawler/go-course/pkg/config"
)

const templatePath = "./templates"

var functions = template.FuncMap{}

// NewTemplateCache allocates a new template cache
func NewTemplateCache(app *config.AppConfig) error {
	myCache := map[string]*template.Template{}

	// public pages
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", templatePath))
	if err != nil {
		return err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", templatePath))
		if err != nil {
			fmt.Println(err)
			return err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", templatePath))
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		matches, err = filepath.Glob(fmt.Sprintf("%s/*.partial.tmpl", templatePath))
		if err != nil {
			fmt.Println(err)
			return err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.partial.tmpl", templatePath))
			if err != nil {
				return err
			}
		}

		matches, err = filepath.Glob(fmt.Sprintf("%s/partials/*.partial.tmpl", templatePath))
		if err != nil {
			return err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/partials/*.partial.tmpl", templatePath))
			if err != nil {
				return err
			}
		}

		// Add the template set to the cache,
		myCache[name] = ts
	}

	app.TemplateCache = myCache

	return nil
}
