package main

import "os"

func main() {
  a := App{}
  a.Initialize(
    os.Getenv("QS_GOLANG_DB_USERNAME"),
    os.Getenv("QS_GOLANG_DB_PASSWORD"),
    os.Getenv("QS_GOLANG_DB_NAME"))

  a.Run(":3000")
}
