package signal

type SigChanType struct {
	Start  chan bool
	Stop   chan bool
	Status chan bool
	Graph  chan bool
	Kill   chan bool
}

var Start = make(chan bool, 1)
var Stop = make(chan bool, 1)
var Status = make(chan bool, 1)
var Graph = make(chan bool, 1)
var Kill = make(chan bool, 1)

var SigChan = SigChanType{
	Start:  Start,
	Stop:   Stop,
	Status: Status,
	Graph:  Graph,
	Kill:   Kill,
}

// func KillAllChan() {
// 	log.Println("Closing all channels")
// 	defer close(Start)
// 	defer close(Stop)
// 	defer close(Status)
// 	defer close(Graph)
// 	defer close(Kill)
// }
