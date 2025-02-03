package backend

import (
	"database/sql"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is GET")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is POST")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is DELETE")
}

func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/", helloWorld)
	a.Router.HandleFunc("/", getRequest).Methods("GET")
	a.Router.HandleFunc("/", postRequest).Methods("POST")
	a.Router.HandleFunc("/", deleteRequest).Methods("DELETE")
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	http.Handle("/", a.Router)
}

func (a *App) RunServer() {
	fmt.Printf("Server started and listening on localhost%s\n", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}
