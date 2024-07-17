package main

import (
	"context"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"github.com/utkarsh-1905/scm/scm"
	"github.com/utkarsh-1905/scm/utils"
)

func main() {
	log.Println("Starting Metrics Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.AddAndParseFlags()
	processes := utils.GetProcsByName()

	influxToken := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, influxToken)
	defer client.Close()
	org := "scm"
	bucket := "scm_monitoring"
	writeAPI := client.WriteAPI(org, bucket)

	ready, err := client.Ready(context.Background())
	if err != nil {
		log.Println("Error in InfluxDB: ", err)
	}
	log.Println("Influx DB Ready? ", *ready.Status, " Since: ", *ready.Up)

	scm.SCM(processes, writeAPI)
}
