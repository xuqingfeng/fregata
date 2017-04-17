test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build-linux-64: fmt
	cd cmd/fregata && GOOS=linux GOARCH=amd64 go build -o ../../out/fregata-linux-amd64 && cd - && cd cmd/fregatad && GOOS=linux GOARCH=amd64 go build -o ../../out/fregatad-linux-amd64

build: fmt
	cd cmd/fregata && go build -o ../../out/fregata && cd - && cd cmd/fregatad && go build -o ../../out/fregatad

run: build
	./out/fregatad -config ./out/fregata.conf
