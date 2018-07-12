
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

func (m *Meal) AddMealFood(db *gorm.DB, foodId string) Output {
  var meal Meal
  var food Food

  if err := db.First(&meal, m.ID).Error; err != nil {
    return err
  }

  if err := db.First(&food, foodID).Error; err != nil {
    return err
  }

  meal_food := MealFood{Food_id: food.ID, Meal_id: meal.ID}
  db.NewRecord(meal_food)
  if err := db.Create(&meal_food).Error; err != nil {
    return err
  }

  output := struct {
    Message string
   }{
     Message: fmt.Sprintf("Successfully added %s to %s", food.Name, meal.Name)
     }
     
  return output
}
