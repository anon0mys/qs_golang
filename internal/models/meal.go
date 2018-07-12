
package models

import (
  "github.com/jinzhu/gorm"
)

type Meal struct {
  ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  Name string `gorm:"not null;" json:"name"`
  Foods []Food `gorm:"many2many:meal_foods;" json:"foods"`
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
