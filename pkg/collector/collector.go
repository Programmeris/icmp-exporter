package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type IcmpCollector struct {
	addressAvailableByIcmp *prometheus.Desc
}

func NewIcmpCollector() *IcmpCollector {
	return &IcmpCollector{
		addressAvailableByIcmp: prometheus.NewDesc("address_available_by_icmp",
			"address availability by icmp",
			nil, nil,
		),
	}
}

func (collector *IcmpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.addressAvailableByIcmp
}

//Collect implements required collect function for all promehteus collectors
func (collector *IcmpCollector) Collect(ch chan<- prometheus.Metric) {
	m1 := prometheus.MustNewConstMetric(collector.addressAvailableByIcmp, prometheus.GaugeValue, 1)
	ch <- m1
}