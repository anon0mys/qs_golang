
package models

import (
  "github.com/jinzhu/gorm"
)

type Food struct {
  ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  Name string `gorm:"not null;"`
  Calories int `gorm:"not null;"`
}

func (f *Food) GetFoods(db *gorm.DB) []Food {
  var foods []Food

  db.Select("ID, Name, Calories").Find(&foods)

  return foods
}

func (f *Food) UpdateFood(db *gorm.DB) error {
  food := Food{}

  if err := db.First(&food, f.ID).Error; err != nil {
    return err
  }

  food.Name = f.Name
  food.Calories = f.Calories

  db.Save(&food)

  return nil
}

func (f *Food) DeleteFood(db *gorm.DB) error {
  if err := db.Delete(&f).Error; err != nil {
    return err
  }
  return nil
}

func (f *Food) CreateFood(db *gorm.DB) error {
  db.NewRecord(&f)

  if err := db.Create(&f).Error; err != nil {
    return err
  }

  return nil
}

func (f *Food) GetFood(db *gorm.DB) error {
  if err := db.Select("id, name, calories").First(&f, f.ID).Error; err != nil {
    return err
  }

  return nil
}
