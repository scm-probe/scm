package utils

import (
	"flag"
	"log"
	"os"
)

var ProcName = ""
var ProcID int
var ParseDump = false
var HELP = flag.Bool("h", false, "Prints the help message")

func AddAndParseFlags() {

	flag.StringVar(&ProcName, "n", "", "Process Name to trace")
	flag.BoolVar(&ParseDump, "d", false, "Parse the Dump file to show sys call names")
	flag.IntVar(&ProcID, "id", -1, "Process ID to trace")

	flag.Parse()

	if ParseDump {
		ProcessDump()
	}

	if ProcName == "" && ProcID == -1 {
		log.Println("Invalid Process : ", ProcName)
		os.Exit(0)
	}

	if *HELP {
		flag.Usage()
		os.Exit(0)
	}

}
