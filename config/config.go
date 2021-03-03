package config

//MetricPort ...
type MetricPort int

var metricPort MetricPort = 9992

//ProvideMetricPort ...
func ProvideMetricPort() MetricPort {
	return metricPort
}
