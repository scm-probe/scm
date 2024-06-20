package exporter

import (
	"github.com/cilium/ebpf"
)

func UpdateMetrics(m *ebpf.Map) {
	defer m.Close()
	itr := m.Iterate()

	var (
		key uint64
		val uint64
	)

	for itr.Next(&key, &val) {
		MetricsParameters[key].Set(float64(val))
	}
}
