package main

import (
  "net/http"
  "os/exec"
  "fmt"
)
type StatusError struct {
  Code int
  Err  error
}

func printWeather(w http.ResponseWriter, r *http.Request) {
  out, err := exec.Command("wego").Output()

  if err != nil {
    w.WriteHeader(500)
  }

  fmt.Fprintf(w, "%s", out)
}

func main() {
  http.HandleFunc("/", printWeather)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
