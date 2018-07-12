package main

import (
  "github.com/anon0mys/qs_golang/api"
)

func main() {
  a := api.App{}
  a.Initialize()

  a.Run()
}
