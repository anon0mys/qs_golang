
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

func (f *Food) GetFoods(db *sql.DB) ([]Food, error) {
  rows, err := db.Query("SELECT id, name, calories FROM foods")

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  foods := []Food{}

  for rows.Next() {
    var f Food
    if err := rows.Scan(&f.ID, &f.Name, &f.Calories); err != nil {
      return nil, err
    }
    foods = append(foods, f)
  }

  return foods, nil
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
  return db.QueryRow("SELECT id, name, calories FROM foods WHERE id=$1",
    f.ID).Scan(&f.ID, &f.Name, &f.Calories)
}
