package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	exp "github.com/utkarsh-1905/scm/exporter"
	"github.com/utkarsh-1905/scm/scm"
	"github.com/utkarsh-1905/scm/utils"
)

func main() {
	log.Println("Starting Metrics Server")

	utils.AddAndParseFlags()
	processes := utils.GetProcsByName()

	exp.MakeMetrics()

	go scm.SCM(processes)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1910", nil)
}
