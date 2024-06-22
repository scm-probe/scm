package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetProcsByName() []int {
	processes, _ := exec.Command("/bin/sh", "-c", fmt.Sprintf("pgrep %s", ProcName)).Output()
	var Procs []int
	if len(processes) == 0 {
		log.Println("Getting Process: No Process Found")
		os.Exit(1)
	}
	proc := strings.Split(string(processes), "\n")
	for _, p := range proc {
		temp, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		Procs = append(Procs, temp)
	}
	return Procs
}
