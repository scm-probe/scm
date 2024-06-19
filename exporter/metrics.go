package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/utkarsh-1905/scm/syscall"
)

var MetricsParameters map[uint64]prometheus.Counter

func MakeMetrics() map[uint64]prometheus.Counter {
	log.Println("Making Metrics")
	MetricsParameters = syscall.ParseSysCallTableToPromCounter()
	return MetricsParameters
}
