package exporter

import (
	"log"
	"time"

	"github.com/cilium/ebpf"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/utkarsh-1905/scm/syscall"
)

var TAGS = map[string]string{
	"scm": "syscall",
}

var SYSCALL_TABLE = syscall.ParseSysCallTableToString()

var EXCLUDE_CALLS = []uint64{228, 96}

func ShouldExcludeCall(call uint64) bool {
	for _, c := range EXCLUDE_CALLS {
		if c == call {
			return true
		}
	}
	return false
}

func UpdateMetrics(m *ebpf.Map, influxWrite api.WriteAPI) {
	itr := m.Iterate()

	var (
		key uint64
		val uint64
	)
	log.Println("Updating Metrics")
	fields := map[string]interface{}{}
	for itr.Next(&key, &val) {
		if ShouldExcludeCall(key) == true {
			continue
		}
		log.Println(key, val)
		name := "scm_" + SYSCALL_TABLE[key]
		fields[name] = val
	}
	point := write.NewPoint("syscall_tracking", TAGS, fields, time.Now())
	influxWrite.WritePoint(point)
}
