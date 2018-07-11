
package api

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "database/sql"

  "github.com/anon0mys/qs_golang/models"

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
  a.initializeRoutes()
}

func (a *App) Run(addr string) {
  log.Fatal(http.ListenAndServe(":3000", a.Router))
}

func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/api/v1/foods", a.CreateFood).Methods("POST")
}

func (a *App) CreateFood(w http.ResponseWriter, r *http.Request) {
  var f models.Food
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&f); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()

  if err := f.CreateFood(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusCreated, f)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
