package main

import (
  "os"
  "github.com/anon0mys/qs_golang/api"
)

func main() {
  dbname := os.Getenv("QS_GOLANG_DB_NAME")
  username := os.Getenv("QS_GOLANG_DB_USERNAME")
  password := os.Getenv("QS_GOLANG_DB_PASSWORD")
  host := os.Getenv("QS_GOLANG_DB_HOST")
  port := os.Getenv("QS_GOLANG_DB_PORT")
  a := api.App{}
  a.Initialize(dbname, username, password, host, port)

  a.Run()
}
