generate:
	cd scm && go generate

run:
	go run main.go

build:
	go build main.go

exec:
	sudo ./main

syscall-table:
	ausyscall $(uname -r) --dump > syscall.csv

provision:
	docker compose -f monitor.docker-compose.yaml up -d

tear:
	docker compose -f monitor.docker-compose.yaml down

clean-influx:
	docker exec -it influxdb influx delete --bucket scm_monitoring \
  --start 2024-03-01T00:00:00Z \
  --stop 2024-11-14T00:00:00Z

graph:
	dot -Tsvg -O temp/graph.gv
	echo "Graph svg is created in temp folder"
	# for GNOME users
	xdg-open temp/graph.gv.svg
