package utils

import (
	"flag"
	"log"
	"os"
)

var ProcName = ""
var ParseDump = false
var HELP = flag.Bool("h", false, "Prints the help message")

func AddAndParseFlags() {

	flag.StringVar(&ProcName, "n", "", "Process ID to trace")
	flag.BoolVar(&ParseDump, "d", false, "Parse the Dump file to show sys call names")

	flag.Parse()

	if ProcName == "" {
		log.Println("Invalid Proc Name: ", ProcName)
		os.Exit(0)
	}

	if ParseDump {
		ProcessDump()
	}

	if *HELP {
		flag.Usage()
		os.Exit(0)
	}

}
