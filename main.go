package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/monofuel/webapp-template/api"
	. "github.com/monofuel/webapp-template/util"
)

var webPort = os.Getenv("PORT")
var templates = template.Must(template.ParseFiles("public/index.html"))

func init() {
	if webPort == "" {
		webPort = "8085"
	}
	//TODO connect to database
}

func main() {
	fmt.Println("starting example webapp")
	registerHandlers()

	log.Printf("listening on port %s", webPort)
	log.Fatal(http.ListenAndServe(":"+webPort, nil))
}

func registerHandlers() {

	r := mux.NewRouter()

	r.Methods("GET").Path("/").Handler(AppHandler(rootHandler))

	jsFs := http.FileServer(http.Dir("public/js"))
	cssFs := http.FileServer(http.Dir("public/css"))
	r.Methods("GET").PathPrefix("/js/").Handler(http.StripPrefix("/js", jsFs))
	r.Methods("GET").PathPrefix("/css/").Handler(http.StripPrefix("/css", cssFs))

	api.AddRoutes(r)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) *AppError {

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return NewAppError(err, "failed to load index.html", 500)
	}
	return nil
}
