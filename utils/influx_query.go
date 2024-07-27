package utils

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxQueryResult struct {
	Time  time.Time   `json:"time"`
	Value interface{} `json:"value"`
}

var DerivativeQuery = `from(bucket: "scm_monitoring")
 |> range(start: -30m)
 |> filter(fn: (r) => r["_measurement"] == "syscall_tracking")
 |> derivative(unit: 1s, nonNegative: false)
 |> yield(name: "derivative")`

var DoubleDerivative = `from(bucket: "scm_monitoring")
   |> range(start: -30m)
   |> filter(fn: (r) => r["_measurement"] == "syscall_tracking")
   |> derivative(unit: 1s, nonNegative: false)
   |> derivative(unit: 1s, nonNegative: false)
   |> yield(name: "double_derivative")`

var Increase = `from(bucket: "scm_monitoring")
   |> range(start: -30m)
   |> filter(fn: (r) => r["_measurement"] == "syscall_tracking")
   |> increase()
   |> yield(name: "increase")`

var EMA = `from(bucket: "scm_monitoring")
   |> range(start: -30m)
   |> filter(fn: (r) => r["_measurement"] == "syscall_tracking")
   |> exponentialMovingAverage(n: 5)
   |> yield(name: "mean")`

func QueryInfluxDB(query string) (string, error) {
	var queryString string
	switch query {
	case "derivative":
		queryString = DerivativeQuery
	case "double_derivative":
		queryString = DoubleDerivative
	case "increase":
		queryString = Increase
	case "ema":
		queryString = EMA
	}
	result, err := QueryAPI.QueryRaw(context.Background(), queryString, influxdb2.DefaultDialect())

	if err != nil {
		log.Println("Error in Querying InfluxDB: ", err)
		return "", err
	}

	return result, nil

}
