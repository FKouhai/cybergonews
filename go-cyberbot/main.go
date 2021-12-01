package main

import (
  "fmt"
  "go-cyberbot/bot"
  "go-cyberbot/config"
)
func main() {
  err := config.ReadConfig()
  if err != nil {
    fmt.Println("err.Error()")
    return
  }
  bot.Start()
  <-make(chan struct {})
  return
}


