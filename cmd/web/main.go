package main

import (
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD

	"github.com/mrpuurple/go-hello-world-web/pkg/config"
	"github.com/mrpuurple/go-hello-world-web/pkg/handlers"
	"github.com/mrpuurple/go-hello-world-web/pkg/render"
)

const portNumber = "8080"

func main() {
	var app config.AppConfig
=======
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrpuurple/mygoapp/pkg/config"
	"github.com/mrpuurple/mygoapp/pkg/handlers"
	"github.com/mrpuurple/mygoapp/pkg/render"
)

const portNumber = ":8888"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
>>>>>>> 888a37c (refactor)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
<<<<<<< HEAD

	app.TemplateCache = tc
	app.UseCache = false
=======
	
	app.TemplateCache = tc
	app.UseCache = true
>>>>>>> 888a37c (refactor)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

<<<<<<< HEAD
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting webserver on port %q\n", portNumber)
	_ = http.ListenAndServe(":"+portNumber, nil)
}
=======
	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
 
>>>>>>> 888a37c (refactor)
