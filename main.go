package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
		fmt.Printf("ENV: %s, slaves total: %2.f, connected: %2.f, lost %2.f\n", configuration.Mesos.Hosts[i].URL, metrics.SlavesTotal, metrics.SlavesActive, metrics.SlaveDisconected)
	}
	fmt.Printf("Total slaves: %2.f, connected: %2.f, lost %2.f\n", metSumary.SlavesTotal, metSumary.SlavesActive, metSumary.SlaveDisconected)
}
