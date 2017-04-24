test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build-linux-64: fmt
	GOOS=linux GOARCH=amd64 go build -o out/fregata-linux-amd64 cmd/fregata/main.go && GOOS=linux GOARCH=amd64 go build -o out/fregatad-linux-amd64 cmd/fregatad/main.go

build: fmt
	go build -o out/fregata cmd/fregata/main.go && go build -o ../../out/fregatad cmd/fregatad/main.go

run: build
	./out/fregatad -config ./out/fregata.conf
