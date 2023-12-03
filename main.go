package main

import (
	. "icmp-exporter/pkg/collector"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    icmpCollector := NewIcmpCollector()
    prometheus.MustRegister(icmpCollector)

    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}