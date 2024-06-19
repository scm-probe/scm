package exporter

import (
	"log"

	"github.com/cilium/ebpf"
)

func UpdateMetrics(m *ebpf.Map) {
	defer m.Close()
	log.Println("Updating Metrics")
	itr := m.Iterate()

	var (
		key uint64
		val uint64
	)

	for itr.Next(&key, &val) {
		MetricsParameters[key].Add(float64(val))
	}
}
