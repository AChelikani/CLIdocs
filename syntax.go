package main

import (
  "net/http"
  "encoding/json"
)

func getJSON(url string, target interface{}) error {
  resp, error := http.Get("http://jservice.io/api/random")
  if error != nil {
    // Deal with error
  }
  defer resp.Body.Close()

  return json.NewDecoder(resp.Body).Decode(target)
}
