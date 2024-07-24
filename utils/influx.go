package utils

import (
	"context"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func StartInfluxDB() api.WriteAPIBlocking {
	influxToken := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, influxToken)
	defer client.Close()
	org := "scm"
	bucket := "scm_monitoring"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	ready, err := client.Ready(context.Background())
	if err != nil {
		log.Println("Error in InfluxDB: ", err)
	}
	log.Println("Influx DB Ready? ", *ready.Status, " Since: ", *ready.Up)

	return writeAPI
}
