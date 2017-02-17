# smetrics
Get metrics from Mesos and K8s.
#### Configuration
Use JSON to pass all **Mesos** and **K8s** environments to get metrics.

##### JSON Example
```json
{
    "mesos": {
        "hosts": [
            {
              "url": "mesos-master-1.domain.pl",
              "port": "8000"
            },
            {
              "url": "mesos-master-2.domain.pl",
              "port": "8000"
            },
            {
              "url": "mesos-master-3.domain.pl",
              "port": "8000"
            }
        ]
    },
    "k8s": {
        "hosts": [
            {
              "url":"k8s-master-1.domain.pl",
              "port": "8999"
            },
            {
              "url":"k8s-master-2.domain.pl",
              "port": "8999"
            }
        ]
    }
}
```
