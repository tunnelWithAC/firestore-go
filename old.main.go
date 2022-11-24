package main

import (
  "context"
  "firebase.google.com/go"
  "google.golang.org/api/option"
)

func main() {

  conf := &firebase.Config{ProjectID: projectID}
  app, err := firebase.NewApp(context.Background(), conf)

  defer client.Close()
}

func getQuote() *Quote {
  "quote": "Play it",
  "author": "Casablanca"
}