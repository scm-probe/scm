package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetProcsByName() ([]int, error) {
	if ProcID > 0 {
		return []int{ProcID}, nil
	}
	processes, _ := exec.Command("/bin/sh", "-c", fmt.Sprintf("pgrep %s", ProcName)).Output()
	var Procs []int
	if len(processes) == 0 {
		log.Println("Getting Process: No Process Found")
		return nil, fmt.Errorf("No Process Found")
	}
	proc := strings.Split(string(processes), "\n")
	for _, p := range proc {
		temp, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		Procs = append(Procs, temp)
	}
	return Procs, nil
}
