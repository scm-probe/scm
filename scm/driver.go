package scm

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/utkarsh-1905/scm/exporter"
	sc_graph "github.com/utkarsh-1905/scm/graph"
	"github.com/utkarsh-1905/scm/signal"
	"github.com/utkarsh-1905/scm/utils"
)

// SCM: system-call-monitor
func SCM(procIDs []int, influxWrite api.WriteAPIBlocking) {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	var objs scmObjects
	if err := loadScmObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// setting the pid of this process
	for _, p := range procIDs {
		log.Println("Running process id: ", p)
		err := objs.ProcMap.Put(uint32(p), uint16(1))
		if err != nil {
			log.Println("Putting Process in Map: ", err)
			os.Exit(1)
		}
	}

	rtp, err := link.AttachRawTracepoint(link.RawTracepointOptions{
		Program: objs.BpfProg,
		Name:    "sys_enter",
	})

	if err != nil {
		log.Println("Attach Tracepoint: ", err)
	}

	sysExitFork, err := link.Tracepoint("syscalls", "sys_exit_clone", objs.AddClone, nil)

	if err != nil {
		log.Println("Attach Tracepoint: ", err)
	}

	defer rtp.Close()
	defer sysExitFork.Close()

	if err != nil {
		log.Println("Putting Process ID: ", err)
	}

	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()

	if utils.Graph {
		go sc_graph.ReadQueue(objs.CallQueue)
	}

	ctx := context.Background()
	for {
		select {
		case <-tick.C:
			go exporter.UpdateMetrics(objs.SysCalls, ctx, influxWrite)
		case <-signal.SigChan.Stop:
			sc_graph.DrawGraph()
			log.Println("Received signal, exiting..")
			return
		}
	}
}
