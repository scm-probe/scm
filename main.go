package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	exp "github.com/utkarsh-1905/scm/exporter"
	"github.com/utkarsh-1905/scm/scm"
	"github.com/utkarsh-1905/scm/utils"
)

func main() {
	log.Println("Starting Metrics Server")

	var ProcName = ""
	flag.StringVar(&ProcName, "n", "", "Process ID to trace")
	flag.Parse()

	if ProcName == "" {
		log.Println("Invalid Proc Name: ", ProcName)
		os.Exit(0)
	}

	processes := utils.GetProcsByName(ProcName)

	exp.MakeMetrics()

	go scm.SCM(processes)
	// utils.ProcessDump()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1910", nil)
}
