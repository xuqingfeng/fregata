test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build-linux-64: fmt
	cd cmd/fregatad && GOOS=linux GOARCH=amd64 go build -o ../../out/fregatad-linux-amd64

build: fmt
	cd cmd/fregatad && go build -o ../../out/fregatad

run: build
	./out/fregatad -config ./out/fregata.conf
