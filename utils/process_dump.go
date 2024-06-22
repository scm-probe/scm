package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/utkarsh-1905/scm/syscall"
)

type DumpUint struct {
	Key   uint64 `json:"key"`
	Value uint64 `json:"value"`
}

type DumpString struct {
	Key   string `json:"key"`
	Value uint64 `json:"value"`
}

func ProcessDump() {
	f, err := os.Open("dump.json")

	if err != nil {
		log.Println("Opening Dump File: ", err)
	}
	defer f.Close()

	j, err := io.ReadAll(f)

	if err != nil {
		log.Println("Reading Dump File: ", err)
	}

	var vals []DumpUint
	var output []DumpString

	json.Unmarshal(j, &vals)

	table := syscall.ParseSysCallTableToString()

	for _, v := range vals {
		var o DumpString
		o.Key = table[int(v.Key)]
		o.Value = v.Value
		output = append(output, o)
	}

	opt, err := json.MarshalIndent(output, "", "	")
	if err != nil {
		log.Println("Formatting output: ", err)
	}

	err = os.WriteFile("dump_processed.json", opt, 0644)

	if err != nil {
		log.Println("Writing output: ", err)
	}
}
