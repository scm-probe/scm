package sc_graph

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/cilium/ebpf"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/utkarsh-1905/scm/syscall"
)

var SYSCALL_TABLE = syscall.ParseSysCallTableToString()

func ReadQueue(queue *ebpf.Map) {
	defer queue.Close()
	var currValue uint64
	prevValue := "START"
	AddVertex(prevValue)
	tick := time.NewTicker(time.Millisecond * 5) //needs tuning according to sys freq
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			err := queue.LookupAndDelete(nil, &currValue)
			if err == nil {
				currCall := SYSCALL_TABLE[currValue]
				AddVertex(currCall)
				AddEdge(prevValue, currCall)
				prevValue = currCall
			}
		}
	}
}

var G = graph.New(graph.StringHash, graph.Acyclic(), graph.Directed(), graph.Weighted())

func AddVertex(v string) {
	_, err := G.Vertex(v)
	if err != nil {
		err := G.AddVertex(v)
		if err != nil {
			log.Println("Error adding vertex: ", err)
		}
	}
}

func AddEdge(from, to string) {
	edge, err := G.Edge(from, to)
	if err != nil {
		err := G.AddEdge(from, to, graph.EdgeWeight(1))
		if err != nil {
			log.Println("Error adding edge: ", err)
		}
	} else {
		weight := ComputeWeight(edge.Properties.Weight)
		err := G.UpdateEdge(from, to, graph.EdgeWeight(weight))
		if err != nil {
			log.Println("Error updating edge: ", err)
		}
	}
}

func ComputeWeight(weight int) int {
	return weight + 1
}

func DrawGraph() {
	order, _ := G.Order()
	size, _ := G.Size()
	log.Println("*************************")
	log.Println("Order of graph is: ", order)
	log.Println("Size of graph is: ", size)
	log.Println("*************************")
	file, _ := os.Create("dot/graph.gv")
	_ = draw.DOT(G, file, draw.GraphAttribute("label", "System Call Sequence"))
}

func DrawGraphOutputIO(w *bytes.Buffer) {
	_ = draw.DOT(G, w, draw.GraphAttribute("label", "System Call Sequence"))
}
