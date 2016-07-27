package main

import (
  "net/http"
  //"log"
  "io/ioutil"
)

func getJSON(url string) []byte {
  resp, err := http.Get("http://localhost:8080")
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
