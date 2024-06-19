generate:
	cd scm && go generate

run:
	go run main.go

build:
	go build main.go

exec:
	sudo ./main
