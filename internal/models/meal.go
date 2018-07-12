
package models

import (
  "github.com/jinzhu/gorm"
)

type Meal struct {
  gorm.Model
  ID int `gorm:PRIMARY_KEY`
  Name string `gorm:"not null;"`
  Foods []Food `gorm:"many2many:meal_foods;"`
}

func (m *Meal) GetMeals(db *gorm.DB) []Meal {
  var meals []Meal

  db.Preload("Foods").Find(&meals)

  return meals
}
