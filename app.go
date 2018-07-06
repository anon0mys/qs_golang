
package main

import (
  "fmt"
  "log"
  "database/sql"
  "net/http"

  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

type App struct {
  Router *mux.Router
  DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
  connectionString :=
    fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)

  var err error
  a.DB, err = sql.Open("postgres", connectionString)
  if err != nil {
    log.Fatal(err)
  }

  a.Router = mux.NewRouter()
  a.Router.HandleFunc("/", getHome).Methods("GET")
  log.Fatal(http.ListenAndServe(":3000", a.Router))
}

func (a *App) Run(addr string) { }

func getHome(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Quantified Self: GoLang")
}
