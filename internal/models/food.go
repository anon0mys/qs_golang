
package models

import (
  "database/sql"
  "errors"
)

type Food struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Calories int `json:"calories"`
}

func (f *Food) GetFoods(db *sql.DB, start, count int) ([]Food, error) {
  return nil, errors.New("Not implemented")
}

func (f *Food) UpdateFood(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (f *Food) DeleteFood(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (f *Food) CreateFood(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO foods(name, calories) VALUES($1, $2) RETURNING id, name, calories",
    f.Name, f.Calories).Scan(&f.ID, &f.Name, &f.Calories)

  if err != nil {
    return err
  }

  return nil
}

func (f *Food) GetFood(db *sql.DB) error {
  return errors.New("Not implemented")
}
