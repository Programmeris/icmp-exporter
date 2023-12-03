package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type IcmpCollector struct {
	addressAvailableByIcmp *prometheus.Desc
}