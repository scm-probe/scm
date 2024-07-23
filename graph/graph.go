package sc_graph

import (
	"log"
	"os"
	"time"

	"github.com/cilium/ebpf"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func ReadQueue(queue *ebpf.Map) {
	var currValue uint64
	prevValue := -1
	AddVertex(prevValue)
	tick := time.NewTicker(time.Millisecond * 10) //needs tuning according to sys freq
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			err := queue.LookupAndDelete(nil, &currValue)
			if err == nil {
				AddVertex(int(currValue))
				AddEdge(int(prevValue), int(currValue))
				prevValue = int(currValue)
			}
		}
	}
}

var G = graph.New(graph.IntHash, graph.Directed(), graph.Acyclic())

func AddVertex(v int) {
	_, err := G.Vertex(v)
	if err != nil {
		G.AddVertex(v)
	}
}

func AddEdge(from, to int) {
	_, err := G.Edge(from, to)
	if err != nil {
		G.AddEdge(from, to)
	}
}

func DrawGraph() {
	order, _ := G.Order()
	log.Println("Order of graph is: ", order)
	file, _ := os.Create("temp/graph.gv")
	_ = draw.DOT(G, file)
}
