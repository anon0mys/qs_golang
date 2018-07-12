package api_test

import (
  "testing"
  "log"
  "os"

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
  app.Initialize()

  app.Run()

  if _, err := app.DB.Instance.Exec(foodsTableCreateQuery); err != nil {
      log.Fatal(err)
  }

  app.DB.Instance.Exec("DELETE FROM foods")
  app.DB.Instance.Exec("ALTER SEQUENCE foods_id_seq RESTART WITH 1")
})

var _ = AfterSuite(func() {
  os.Exit(0)
})

const foodsTableCreateQuery = `CREATE TABLE IF NOT EXISTS foods
(
id SERIAL,
name TEXT NOT NULL,
calories INTEGER NOT NULL,
CONSTRAINT foods_pkey PRIMARY KEY (id)
)`

const mealsTableCreateQuery = `CREATE TABLE IF NOT EXISTS meals
(
id SERIAL,
name TEXT NOT NULL,
CONSTRAINT meals_pkey PRIMARY KEY (id)
)`

const mealFoodsTableCreateQuery = `CREATE TABLE IF NOT EXISTS meal_foods
(
food_id INTEGER REFERENCES foods,
meal_id INTEGER REFERENCES meals
)`
