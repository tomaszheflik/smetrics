package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type hosts struct {
	url  string
	port int
}

type cluster struct {
	hosts []hosts
}

type config struct {
	mesos cluster
	k8s   cluster
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
