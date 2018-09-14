https://prometheus.io/docs/prometheus/latest/installation/

https://github.com/prometheus/prometheus/blob/release-2.3/documentation/examples/prometheus.yml
https://github.com/prometheus/prometheus/blob/release-2.3/documentation/examples/prometheus-kubernetes.yml

```
docker run -p 9090:9090 -v ${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml \
       prom/prometheus
```
