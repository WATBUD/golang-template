package main

import (
    "log"
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  good := os.Getenv("good")
  
  
  fmt.Println(good);
  goo := os.Getenv("goo")
  
  fmt.Println(goo);
  
  // now do something with s3 or whatever
}