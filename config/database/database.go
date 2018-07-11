package database

import (
  "fmt"
  "database/sql"
  "log"
  "os"

  _ "github.com/lib/pq"
)

type DB struct {
  Instance *sql.DB
}

func Initialize() *DB {
  user := os.Getenv("QS_GOLANG_DB_USERNAME")
  password := os.Getenv("QS_GOLANG_DB_PASSWORD")
  dbname := os.Getenv("QS_GOLANG_DB_NAME")
  connectionParams := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)

  var err error
  instance, err := sql.Open("postgres", connectionParams)
  if err != nil {
    log.Fatal(err)
  }
  return &DB {Instance: instance}
}
