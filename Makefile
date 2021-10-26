run:
	go run . ./testfiles/run.go

build:
	go build -ldflags "-s -w" .

install: 
	go install .
	