# smetrics
Get metrics from Mesos and K8s.
#### Configuration
Use JSON to pass all **Mesos** and **K8s** environments to get metrics.

##### JSON Example
```json
{
    "mesos": { 
        "hosts": [
            "mesos-master-1", 
            "mesos-master-2", 
            "mesos-master-3", 
            "mesos-master-4"
        ]   
    },
    "k8s": {
        "hosts": [
            "k8s-master-1",
            "k8s-master-2",
            "k8s-master-3"
        ]
    }
}
```
