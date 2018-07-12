
package models

import (
  "database/sql"
)

type Meal struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Foods []Food
}

func (m *Meal) GetMeals(db *sql.DB) ([]Meal, error) {
  rows, err := db.Query(`
    SELECT meals.*,
    COALESCE(json_agg(foods.*) FILTER (WHERE foods.id IS NOT NULL), '[]') AS foods
    FROM meals
    LEFT JOIN meal_foods ON meals.id = meal_foods.meal_id
    LEFT JOIN foods ON meal_foods.food_id = foods.id
    GROUP BY meals.id;
    `)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  meals := []Meal{}

  for rows.Next() {
    var m Meal
    var foods []Food

    if err := rows.Scan(&m.ID, &m.Name, &foods); err != nil {
      return nil, err
    }
    for _, f := range foods {
      m.Foods = append(m.Foods, f)
    }
    meals = append(meals, m)
  }

  return meals, nil
}
