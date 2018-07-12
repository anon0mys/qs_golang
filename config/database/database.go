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
  database := os.Getenv("QS_GOLANG_DB_NAME")
  username := os.Getenv("QS_GOLANG_DB_USERNAME")
  password := os.Getenv("QS_GOLANG_DB_PASSWORD")
  host := os.Getenv("QS_GOLANG_DB_HOST")
  port := os.Getenv("QS_GOLANG_DB_PORT")
  connectionParams := fmt.Sprintf("%s:%s/%s?user=%s&password=%s&sslmode=disable", host, port, database, username, password)

  var err error
  instance, err := sql.Open("postgres", connectionParams)
  if err != nil {
    log.Fatal(err)
  }

  return &DB {Instance: instance}
}
