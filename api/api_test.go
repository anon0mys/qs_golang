package api_test

import (
  "testing"
  "os"
  "log"

  api "github.com/anon0mys/qs_golang/api"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var app *api.App

func TestAPI(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "API Suite")
}

var _ = BeforeSuite(func() {
  app := api.App{}
  app.Initialize(
    os.Getenv("QS_GOLANG_DB_USERNAME"),
    os.Getenv("QS_GOLANG_DB_PASSWORD"),
    os.Getenv("QS_GOLANG_DB_NAME"))

  app.Run(":3000")

  if _, err := app.DB.Exec(tableCreationQuery); err != nil {
      log.Fatal(err)
  }

  app.DB.Exec("DELETE FROM foods")
  app.DB.Exec("ALTER SEQUENCE foods_id_seq RESTART WITH 1")
})

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS foods
(
id SERIAL,
name TEXT NOT NULL,
calories INTEGER NOT NULL,
CONSTRAINT foods_pkey PRIMARY KEY (id)
)`
