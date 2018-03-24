VERSION=`git describe --abbrev=0 --tags`

test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build-all: fmt
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-linux-amd64 cmd/fregata/main.go && \
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-linux-amd64 cmd/fregatad/main.go && \
	GOOS=linux GOARCH=386 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-linux-386 cmd/fregata/main.go && \
	GOOS=linux GOARCH=386 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-linux-386 cmd/fregatad/main.go && \
	GOOS=linux GOARCH=arm go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-linux-arm cmd/fregata/main.go && \
	GOOS=linux GOARCH=arm go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-linux-arm cmd/fregatad/main.go && \
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-darwin-amd64 cmd/fregata/main.go && \
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-darwin-amd64 cmd/fregatad/main.go && \
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-windows-amd64.exe cmd/fregata/main.go && \
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-windows-amd64.exe cmd/fregatad/main.go && \
	GOOS=windows GOARCH=386 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata-windows-386.exe cmd/fregata/main.go && \
	GOOS=windows GOARCH=386 go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad-windows-386.exe cmd/fregatad/main.go

build: fmt
	go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregata cmd/fregata/main.go && \
	go build -ldflags "-w -s -X main.version=${VERSION}" -o out/fregatad cmd/fregatad/main.go

run: build
	./out/fregatad -config ./out/fregata.conf
