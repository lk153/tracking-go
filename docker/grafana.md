# GRAFANA BOARD SETUP

## Query setup 

1. Network IO (Bytes transfer)
```
sum by (direction) (increase(system_network_io[1m])) / 1000000
```
2. System memory usage
```
avg_over_time(system_memory_usage[1m]) / 1000000
```
3. API Latency - TP95
```
histogram_quantile(0.95, rate(http_server_requests_duration_seconds_bucket[1m])) * 1000
```
4. API Throughput
```
increase(http_server_requests_number[1m])
```

## Alert setting
1. Get webhook url from **Slack**
2. Create notification channel on **Grafana**
3. Set Alert rule to specific pannel on **Grafana** (_Alert tab_) 
