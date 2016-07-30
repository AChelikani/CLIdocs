package main

import (
  "net/http"
  //"log"
  "io/ioutil"
)

// Base URL for API
const BaseURL string = "http://localhost:8080/"

func getJSON(url string) []byte {
  resp, err := http.Get(BaseURL + url)
  if err != nil {
    // Deal with error
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    // Deal with error
  }
  defer resp.Body.Close()

  return []byte(body)
}
