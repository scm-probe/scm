package utils

import (
	"context"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var WriteAPI api.WriteAPIBlocking
var QueryAPI api.QueryAPI

func StartInfluxDB() {
	influxToken := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, influxToken)
	defer client.Close()
	org := "scm"
	bucket := "scm_monitoring"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)
	ready, err := client.Ready(context.Background())
	if err != nil {
		log.Println("Error in InfluxDB: ", err)
	}
	log.Println("Influx DB Ready? ", *ready.Status, " Since: ", *ready.Up)

	WriteAPI = writeAPI
	QueryAPI = queryAPI
}
