package config

import (
	"html/template"
	"log"
<<<<<<< HEAD
=======

	"github.com/alexedwards/scs/v2"
>>>>>>> 888a37c (refactor)
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
<<<<<<< HEAD
=======
	InProduction  bool
	Session       *scs.SessionManager
>>>>>>> 888a37c (refactor)
}
