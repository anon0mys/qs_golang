package models

import (
  "github.com/jinzhu/gorm"
)

type MealFood struct {
  ID int
  Food_id int
  Meal_id int
}

func (mf *MealFood) CreateMealFood(db *gorm.DB) error {
  db.NewRecord(&mf)

  if err := db.Create(&mf).Error; err != nil {
    return err
  }

  return nil
}

func (mf *MealFood) GetMealFood(db *gorm.DB, meal_id, food_id int) error {
  if err := db.Where("meal_id = ? AND food_id = ?", meal_id, food_id).First(&mf).Error; err != nil {
    return err
  }

  return nil
}

func (mf *MealFood) DeleteMealFood(db *gorm.DB) error {
  if err := db.Delete(&mf).Error; err != nil {
    return err
  }

  return nil
}
