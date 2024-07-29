package utils

import (
	"flag"
	"os"

	"github.com/utkarsh-1905/scm/signal"
)

var ProcName = ""
var ProcID int
var ParseDump = false
var Graph = false
var HELP = flag.Bool("h", false, "Prints the help message")

func AddAndParseFlags() {

	flag.StringVar(&ProcName, "n", "", "Process Name to trace")
	flag.BoolVar(&ParseDump, "d", false, "Parse the Dump file to show sys call names")
	flag.IntVar(&ProcID, "id", -1, "Process ID to trace")
	flag.BoolVar(&Graph, "g", false, "Enable Graph Mode")

	flag.Parse()

	if ParseDump {
		ProcessDump()
	}

	if *HELP {
		flag.Usage()
		os.Exit(0)
	}

}

func CheckCMDFlags() {
	if ProcName != "" || ProcID != -1 {
		signal.SigChan.Start <- true
	}
}
