package database

import (
  "fmt"
  "log"
  "os"

  "github.com/anon0mys/qs_golang/internal/models"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
  Instance *gorm.DB
}

func Initialize() *DB {
  database := os.Getenv("QS_GOLANG_DB_NAME")
  username := os.Getenv("QS_GOLANG_DB_USERNAME")
  password := os.Getenv("QS_GOLANG_DB_PASSWORD")
  host := os.Getenv("QS_GOLANG_DB_HOST")
  port := os.Getenv("QS_GOLANG_DB_PORT")
  connectionParams := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, database, password)

  var err error
  instance, err := gorm.Open("postgres", connectionParams)

  if err != nil {
    log.Fatal(err)
  }

  instance.Table("foods").CreateTable(&models.Food{})
  instance.Table("meals").CreateTable(&models.Meal{})

  instance.AutoMigrate(&models.Food{})
  instance.AutoMigrate(&models.Meal{})

  return &DB {Instance: instance}
}
