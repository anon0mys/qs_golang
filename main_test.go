package main_test

import (
    "log"
    "os"
    "encoding/json"
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/anon0mys/qs_golang"
)

var a main.App

func TestMain(m *testing.M) {
    a = main.App{}
    a.Initialize(
        os.Getenv("TEST_DB_USERNAME"),
        os.Getenv("TEST_DB_PASSWORD"),
        os.Getenv("TEST_DB_NAME"))

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM foods")
    a.DB.Exec("ALTER SEQUENCE foods_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS foods
(
id SERIAL,
name TEXT NOT NULL,
calories INTEGER NOT NULL,
CONSTRAINT foods_pkey PRIMARY KEY (id)
)`

func TestEmptyTable(t *testing.T) {
  clearTable()

  req, _ := http.NewRequest("GET", "/foods", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  a.Router.ServeHTTP(rr, req)

  return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
}

func TestCreateFoods(t *testing.T) {
  clearTable()

  payload := []byte(`{"name":"Banana","calories":100}`)

  req, _ := http.NewRequest("POST", "/api/v1/foods", bytes.NewBuffer(payload))
  response := executeRequest(req)

  checkResponseCode(t, http.StatusCreated, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["name"] != "Banana" {
    t.Errorf("Expected food name to be 'Banana'. Got '%v'", m["name"])
  }

  if m["calories"] != 100.0 {
    t.Errorf("Expected food calories to be '100'. Got '%v'", m["calories"])
  }

  if m["id"] != 1.0 {
    t.Errorf("Expected food id to be '1'. Got '%v'", m["id"])
  }
}
