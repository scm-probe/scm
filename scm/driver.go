package scm

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
	"github.com/utkarsh-1905/scm/exporter"
)

// SCM: system-call-monitor
func SCM() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	var objs scmObjects
	if err := loadScmObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	link, err := link.AttachRawTracepoint(link.RawTracepointOptions{
		Program: objs.BpfProg,
		Name:    "sys_enter",
	})

	if err != nil {
		log.Println("Attach Tracepoint: ", err)
	}

	defer link.Close()

	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)
	for {
		select {
		case <-tick.C:
			dup, err := objs.SysCalls.Clone()
			if err != nil {
				log.Println("Duplicating Buffer: ", err)
			}
			go exporter.UpdateMetrics(dup)
			log.Println("Tick Tock")
		case <-stop:
			if err := objs.SysCalls.Close(); err != nil {
				log.Println("Closing Map: ", err)
			}
			log.Println("Received signal, exiting..")
			os.Exit(0)
			return
		}
	}
}
