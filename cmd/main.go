package main

import (
  "log"

  "github.com/lightbliss/clean_architecture_blog/internal/app"
)

func main() {
  c := app.NewContext()
  log.Fatal(c.WebServer().Start())
}