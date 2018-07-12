
package models

import (
  "github.com/jinzhu/gorm"
)

type Meal struct {
  ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  Name string `gorm:"not null;"`
  Foods []Food `gorm:"many2many:meal_foods;"`
}

func (m *Meal) GetMeals(db *gorm.DB) []Meal {
  var meals []Meal

  db.Preload("Foods").Find(&meals)

  return meals
}

func (m *Meal) GetMeal(db *gorm.DB) Meal {
  var meal Meal

  db.Preload("Foods").First(&meal, m.ID)

  return meal
}
