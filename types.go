package main

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

// MESMetrics holds information on slaves
type MESMetrics struct {
	SlavesTotal      float64
	SlavesActive     float64
	SlaveDisconected float64
}

//SLSMetrics holds information about slave ram, cpu, disks
type SLAMetrics struct {
}
