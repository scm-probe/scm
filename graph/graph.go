package sc_graph

import (
	"log"

	"github.com/cilium/ebpf"
)

var QueueChan chan uint64

func ReadQueue(queue *ebpf.Map) {
	var value interface{}
	err := queue.LookupAndDelete(0, &value)
	log.Println("Queue Value: ", value, "Err: ", err)
	for err != nil {
		log.Println("Queue Value: ", value, "Err: ", err)
		err = queue.LookupAndDelete(0, &value)
	}
}
