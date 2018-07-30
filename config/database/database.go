package database

import (
  "fmt"
  "log"

  "github.com/anon0mys/qs_golang/internal/models"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
  Instance *gorm.DB
}

func Initialize(dbname, username, password, host, port string) *DB {
  connectionParams := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)

  var err error
  instance, err := gorm.Open("postgres", connectionParams)

  if err != nil {
    log.Fatal(err)
  }

  if instance.Table("foods") == nil && instance.Table("meals") == nil {
    instance.Table("foods").CreateTable(&models.Food{})
    instance.Table("meals").CreateTable(&models.Meal{})
  }

  instance.AutoMigrate(&models.Food{})
  instance.AutoMigrate(&models.Meal{})
  instance.AutoMigrate(&models.MealFood{})

  return &DB {Instance: instance}
}
