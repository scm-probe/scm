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
	defer queue.Close()
	var currValue uint64
	prevValue := -1
	AddVertex(prevValue)
	tick := time.NewTicker(time.Millisecond * 5) //needs tuning according to sys freq
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

var G = graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.Weighted())

func AddVertex(v int) {
	_, err := G.Vertex(v)
	if err != nil {
		G.AddVertex(v)
	}
}

func AddEdge(from, to int) {
	edge, err := G.Edge(from, to)
	if err != nil {
		G.AddEdge(from, to, graph.EdgeWeight(1))
	} else {
		weight := ComputeWeight(edge.Properties.Weight)
		G.UpdateEdge(from, to, graph.EdgeWeight(weight))
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
	file, _ := os.Create("temp/graph.gv")
	_ = draw.DOT(G, file, draw.GraphAttribute("label", "System Call Sequence"))
}
