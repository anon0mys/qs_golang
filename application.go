package main

import (
  "os"
  "github.com/anon0mys/qs_golang/api"
)

func main() {
  a := api.App{}
  a.Initialize(
    os.Getenv("QS_GOLANG_DB_USERNAME"),
    os.Getenv("QS_GOLANG_DB_PASSWORD"),
    os.Getenv("QS_GOLANG_DB_NAME"))

  a.Run()
}
