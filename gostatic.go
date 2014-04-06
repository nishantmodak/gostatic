package main

import (
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("public"))
  http.ListenAndServe(":8080", fs)
}