package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/andygrunwald/megos"
)

// Hosts holds cluster definition
type Hosts struct {
	URL      string `json:"url"`
	PORT     string `json:"port"`
	ENDPOINT string `json:"endpoint"`
}

// Cluster definition
type Cluster struct {
	Hosts []Hosts `json:"hosts"`
}

//Config holds all JSON with configurationd
type Config struct {
	Mesos Cluster `json:"mesos"`
	K8s   Cluster `json:"k8s"`
}

func main() {
	content, _ := ioutil.ReadFile("config.json")
	var configuration Config
	err := json.Unmarshal(content, &configuration)
	if err != nil {
		fmt.Println("error-f:", err)
	}
	// Dump(configuration.Mesos.Hosts[0])
	mesosurl := fmt.Sprintf("http://%s:%s/", configuration.Mesos.Hosts[0].URL, configuration.Mesos.Hosts[0].PORT)
	node, _ := url.Parse(mesosurl)
	mesos := megos.NewClient([]*url.URL{node}, nil)
	state, _ := mesos.DetermineLeader()
	metrics, err := mesos.GetMetricsSnapshot(state)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mesos: %s slaves total: %2.f, active: %2.f, disconected: %2.f\n", configuration.Mesos.Hosts[0].URL, metrics.MasterSlavesActive, metrics.MasterSlavesConnected, metrics.MasterSlavesInactive)
	// Dump(metrics.MasterSlavesActive)
}

// Dump proints propperly formated JSON
func Dump(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(result))
}
