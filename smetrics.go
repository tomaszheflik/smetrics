package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type hosts struct {
	url  []string
	port int
}

type mesos struct {
	hosts []hosts
}
type k8s struct {
	hosts []hosts
}

type config struct {
	mesos mesos
	k8s   k8s
}

func main() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error-f:", err)
	}
	fmt.Println(configuration.mesos) // output: [UserA, UserB]
}
