generate:
	cd scm && go generate

run:
	go run main.go

build:
	go build main.go

exec:
	sudo ./main

clean-influx:
	docker exec -it influxdb influx delete --bucket scm_monitoring \
  --start 2024-03-01T00:00:00Z \
  --stop 2024-11-14T00:00:00Z
