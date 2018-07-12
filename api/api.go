
package api

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"

  "github.com/anon0mys/qs_golang/internal/models"
  "github.com/anon0mys/qs_golang/config/database"

  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

type App struct {
  Router *mux.Router
  DB     *database.DB
  Server *http.Server
}

func (a *App) Initialize() {
  a.DB = database.Initialize()
  a.Router = mux.NewRouter()
  a.initializeRoutes()

  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

  a.Server = &http.Server{Addr: ":3000", Handler: handlers.CORS(headersOk, originsOk, methodsOk)(a.Router)}
}

func (a *App) Run() {
  log.Fatal(a.Server.ListenAndServe())
}

func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/api/v1/foods/", a.CreateFood).Methods("POST")
  a.Router.HandleFunc("/api/v1/foods/", a.GetFoods).Methods("GET")
  a.Router.HandleFunc("/api/v1/foods/{id:[0-9]+}", a.GetFood).Methods("GET")
  a.Router.HandleFunc("/api/v1/foods/{id:[0-9]+}", a.UpdateFood).Methods("PUT", "PATCH", "OPTIONS")
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

func (a *App) GetFoods(w http.ResponseWriter, r *http.Request) {
  var f models.Food
  foods, err := f.GetFoods(a.DB.Instance)

  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, foods)
}

func (a *App) GetFood(w http.ResponseWriter, r *http.Request) {
  var f models.Food
  params := mux.Vars(r)

  id, err := strconv.Atoi(params["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid food ID")
    return
  }

  f = models.Food {ID: id}
  if err := f.GetFood(a.DB.Instance); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid product ID")
    return
  }

  respondWithJSON(w, http.StatusOK, f)
}

func (a *App) UpdateFood(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid food ID")
  }

  var f models.Food
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&f); err != nil {
    respondWithError(w, 400, "Food not updated")
    return
  }
  defer r.Body.Close()
  f.ID = id

  if err := f.UpdateFood(a.DB.Instance); err != nil {
    respondWithError(w, 400, "Food not updated")
    return
  }

  respondWithJSON(w, http.StatusOK, f)
}

func (a *App) CreateFood(w http.ResponseWriter, r *http.Request) {
  var f models.Food
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&f); err != nil {
    respondWithError(w, http.StatusBadRequest, "Food not created")
    return
  }
  defer r.Body.Close()

  if err := f.CreateFood(a.DB.Instance); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusCreated, f)
}
