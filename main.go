package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	exp "github.com/utkarsh-1905/scm/exporter"
	"github.com/utkarsh-1905/scm/scm"
)

func main() {
	log.Println("Starting Metrics Server")

	var ScrapeTime = 2
	flag.IntVar(&ScrapeTime, "i", 2, "Scrape Internal from eBPF Buffer")
	flag.Parse()

	exp.MakeMetrics()
	go scm.SCM()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1910", nil)
}
