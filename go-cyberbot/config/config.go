package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

var (
  Token string
  BotPrefix string
  config *configStruct
)

type configStruct struct { //config struct which contains the bot Token and its Prefix
  Token string `json : "Token"`
  BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error { //Function that reads the config file using json unmarshal
  fmt.Println("Reading config file")
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    fmt.Println(err.Error())
    return err
  }
  fmt.Println(string(file))
  err = json.Unmarshal(file, &config)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }
  Token = config.Token
  BotPrefix = config.BotPrefix
  return nil
}
