package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/utkarsh-1905/scm/scm"
	"github.com/utkarsh-1905/scm/server"
	"github.com/utkarsh-1905/scm/signal"
	"github.com/utkarsh-1905/scm/utils"
)

func main() {
	log.Println("Starting Metrics Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.AddAndParseFlags()

	r := mux.NewRouter()

	subRouter := r.PathPrefix("/").Subrouter()
	server.RegisterRoutesAndMiddleware(subRouter)
	r.PathPrefix("/").Handler(subRouter)

	utils.StartInfluxDB()
	utils.CheckCMDFlags()

	go func() {
		for {
			log.Println("Waiting for Signal")
			select {
			case <-signal.SigChan.Start:
				processes, err := utils.GetProcsByName()
				if err != nil {
					log.Println("Error getting processes by name")
					continue
				}
				scm.SCM(processes, utils.WriteAPI)
			case <-signal.SigChan.Kill:
				log.Println("Killing SCM Server")
			}
		}
	}()

	http.ListenAndServe(":1910", r)
	log.Println("SCM Server Started")
}

// Routes Needed
// 1. /start
// 2. /stop
// 4. /status
// 5. /metrics
// 6. /graph
