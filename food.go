
package main

import (
  "database/sql"
  "errors"
)

type food struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Calories int `json:"calories"`
}

func (f *food) getFoods(db *sql.DB, start, count int) ([]food, error) {
  return nil, errors.New("Not implemented")
}

func (f *food) updateFood(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (f *food) deleteFood(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (f *food) createFood(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO foods(name, calories) VALUES($1, $2) RETURNING id, name, calories",
    f.Name, f.Calories).Scan(&f.ID)

  if err != nil {
    return err
  }

  return nil
}

func (f *food) getFood(db *sql.DB) error {
  return errors.New("Not implemented")
}
