package exporter

import (
	"log"

	"github.com/cilium/ebpf"
)

func UpdateMetrics(m *ebpf.Map) {
	itr := m.Iterate()

	var (
		key uint64
		val uint64
	)
	log.Println("Updating Metrics")
	for itr.Next(&key, &val) {
		log.Println(key, val)
		MetricsParameters[key].Set(float64(val))
	}
}
