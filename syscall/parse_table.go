package syscall

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func parseSysCallTable(name string) [][]string {
	file, err := os.Open("syscalls.csv")

	if err != nil {
		log.Println("Error in opening csv table: ", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Println("Error in reading csv: ", err)
	}

	log.Println("Reading: ", len(records), " records")
	return records
}

func ParseSysCallTableToString() map[int]string {

	records := parseSysCallTable("syscall.csv")

	syscall := make(map[int]string)

	for _, record := range records {
		call := record[0]
		args := strings.Split(call, "	")
		if len(args) != 2 {
			continue
		}
		key, _ := strconv.Atoi(args[0])
		syscall[key] = args[1]
	}

	return syscall
}

func ParseSysCallTableToPromCounter() map[uint64]prometheus.Gauge {

	records := parseSysCallTable("syscall.csv")

	syscall := make(map[uint64]prometheus.Gauge)

	for _, record := range records {
		call := record[0]
		args := strings.Split(call, "	")
		if len(args) != 2 {
			continue
		}
		key, _ := strconv.Atoi(args[0])
		syscall[uint64(key)] = promauto.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("scm_%s_call_count", args[1]),
			Help: fmt.Sprintf("Total number of calls for system call: %s", args[1]),
		})
	}

	return syscall
}
