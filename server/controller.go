package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	sc_graph "github.com/utkarsh-1905/scm/graph"
	"github.com/utkarsh-1905/scm/signal"
	"github.com/utkarsh-1905/scm/utils"
)

type StartRequest struct {
	Data     string `json:"data"`
	DataType string `json:"dataType"`
	Graph    bool   `json:"graph"`
}

func Start(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	var req StartRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	if req.DataType == "id" {
		d, err := strconv.Atoi(req.Data)
		if err != nil {
			http.Error(w, "Error parsing request data", http.StatusBadRequest)
			return
		}
		utils.ProcID = d
	} else if req.DataType == "name" {
		utils.ProcName = req.Data
	} else {
		http.Error(w, "Invalid Data Type", http.StatusBadRequest)
		return
	}

	if req.Graph {
		utils.Graph = true
	}

	signal.SigChan.Start <- true
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Starting Process"))
}

func Stop(w http.ResponseWriter, _ *http.Request) {
	signal.SigChan.Stop <- true
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stopping Process"))
}

func Status(w http.ResponseWriter, r *http.Request) {}

func Metrics(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query not provided", http.StatusBadRequest)
		return
	}
	res, err := utils.QueryInfluxDB(query)
	if err != nil {
		http.Error(w, "Error querying InfluxDB", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func Graph(w http.ResponseWriter, r *http.Request) {
	var dot bytes.Buffer
	sc_graph.DrawGraphOutputIO(&dot)
	w.WriteHeader(http.StatusOK)
	w.Write(dot.Bytes())
}
