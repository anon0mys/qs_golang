package api_test

import (
  "testing"
  "os"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "bytes"

  api "github.com/anon0mys/qs_golang/api"
)

var app api.App

func TestMain(m *testing.M) {
  app = api.App{}
  app.Initialize(
    os.Getenv("QS_GOLANG_TEST_DB_NAME"),
    os.Getenv("QS_GOLANG_DB_USERNAME"),
    os.Getenv("QS_GOLANG_DB_PASSWORD"),
    os.Getenv("QS_GOLANG_DB_HOST"),
    os.Getenv("QS_GOLANG_DB_PORT"))

  code := m.Run()

  clearTables()

  os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
  clearTables()

  req, _ := http.NewRequest("GET", "/api/v1/foods", nil)
  response := executeRequest(req)

  checkResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestCreateFood(t *testing.T) {
  clearTables()

  payload := []byte(`{"food":{"name":"Test food","calories":"100"}}`)

  req, _ := http.NewRequest("POST", "/api/v1/foods", bytes.NewBuffer(payload))
  response := executeRequest(req)

  checkResponseCode(t, http.StatusCreated, response.Code)

  var f map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &f)

  if f["name"] != "Test food" {
    t.Errorf("Expected food name to be 'Test food'. Got '%v'", f["name"])
  }

  if f["calories"] != "100" {
    t.Errorf("Expected food calories to be '100'. Got '%v'", f["calories"])
  }

  if f["id"] != 1.0 {
    t.Errorf("Expected food ID to be '1'. Got '%v'", f["id"])
  }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  app.Router.ServeHTTP(rr, req)

  return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
}

func clearTables() {
  app.DB.Instance.Exec("DELETE FROM meal_foods")
  app.DB.Instance.Exec("DELETE FROM foods")
  app.DB.Instance.Exec("ALTER SEQUENCE foods_id_seq RESTART WITH 1")
  app.DB.Instance.Exec("DELETE FROM meals")
  app.DB.Instance.Exec("ALTER SEQUENCE meals_id_seq RESTART WITH 1")
}
