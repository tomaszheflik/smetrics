package main

import (
    "encoding/json"
    "os"
    "fmt"
)

type Host struct {
  url       []string
  port      int
}

type Mesos struct {
    hosts   []Host
}
type K8s struct {
    hosts   []Host
}

type Configuration struct {
  Mesos   Mesos
  K8s     K8s
}


file, _ := os.Open("conf.json")
decoder := json.NewDecoder(file)
configuration := Configuration{}
err := decoder.Decode(&configuration)
if err != nil {
  fmt.Println("error:", err)
}
fmt.Println(configuration.Users) // output: [UserA, UserB]
