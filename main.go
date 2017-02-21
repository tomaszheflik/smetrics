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
	Port     string `json:"port"`
	Endpoint string `json:"endpoint"`
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

// MesMetrics holds information on slaves, ram , cpu
type MesMetrics struct {
	SlavesTotal      float64
	SlavesActive     float64
	SlaveDisconected float64
}

func main() {
	var configuration Config
	var metSumary MesMetrics
	content, _ := ioutil.ReadFile("config.json")
	err := json.Unmarshal(content, &configuration)
	if err != nil {
		fmt.Println("error-f:", err)
	}
	for i := 0; i < len(configuration.Mesos.Hosts); i++ {
		metrics := MesosMetrics(configuration.Mesos.Hosts[i].URL, configuration.Mesos.Hosts[i].Port)
		metSumary.SlavesTotal += metrics.SlavesTotal
		metSumary.SlavesActive += metrics.SlavesActive
		metSumary.SlaveDisconected += metrics.SlaveDisconected
		fmt.Printf("ENV: %s, slaves total: %2.f, slaves connected: %2.f, slave lost %2.f\n", configuration.Mesos.Hosts[i].URL, metrics.SlavesTotal, metrics.SlavesActive, metrics.SlaveDisconected)
	}
	fmt.Printf("Total slaves: %2.f, total slaves connected: %2.f, total slave lost %2.f\n", metSumary.SlavesTotal, metSumary.SlavesActive, metSumary.SlaveDisconected)
}

// Dump proints propperly formated JSON
func Dump(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(result))
}

// MesosMetrics return metrics from particular mesos master
func MesosMetrics(URL string, Port string) (metrics MesMetrics) {
	mesosurl := fmt.Sprintf("http://%s:%s/", URL, Port)
	node, _ := url.Parse(mesosurl)
	mesos := megos.NewClient([]*url.URL{node}, nil)
	state, _ := mesos.DetermineLeader()
	met, err := mesos.GetMetricsSnapshot(state)
	if err != nil {
		panic(err)
	}
	metrics.SlavesTotal = met.MasterSlavesActive
	metrics.SlavesActive = met.MasterSlavesConnected
	metrics.SlaveDisconected = met.MasterSlavesInactive
	return
}
