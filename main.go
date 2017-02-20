package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Hosts struct {
	URL  string `json:"url"`
	PORT string `json:"port"`
}

type Cluster struct {
	Hosts []Hosts `json:"hosts"`
}

type Config struct {
	Mesos Cluster `json:"mesos"`
	K8s   Cluster `json:"k8s"`
}

func main() {
	// file, _ := os.Open("conf.json")
	content, _ := ioutil.ReadFile("config.json")
	var configuration Config
	// fmt.Printf("%s", content)
	err := json.Unmarshal(content, &configuration)
	if err != nil {
		fmt.Println("error-f:", err)
	}
	Dump(len(configuration.K8s.Hosts)) // output: [UserA, UserB

}
func Dump(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(result))
}
