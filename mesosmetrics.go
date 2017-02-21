package main

import (
	"fmt"
	"net/url"

	"github.com/andygrunwald/megos"
)

// MesosMetrics return metrics from particular mesos master
func MesosMetrics(URL string, Port string) (metrics MESMetrics) {
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

//MSlavesMetrics collect metrics from all slaves
func MSlavesMetrics(string string) {

}
