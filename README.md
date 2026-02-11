# fregata
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/xuqingfeng/fregata)

### Install

Download from [release](https://github.com/xuqingfeng/fregata/releases) page

Or install through go cmd:

```bash
# Install the binary to $GOPATH/bin (or $HOME/go/bin)
go install github.com/xuqingfeng/fregata/cmd/fregata@latest
```

Or build from source:

```bash
git clone https://github.com/xuqingfeng/fregata.git && cd fregata
make build
# binary: ./out/fregata
```

### Supported Services

[examples](./examples)

- Slack
- Telegram
- Twilio
- Wechat
- SMTP Email
- MACOS

### Usage

- [fregata.conf](./etc/fregata.conf)
- `fregata -config fregata.conf`

### Docker

```bash
docker run --rm -p 2017:2017 -v `pwd`/fregata.conf:/etc/fregata/fregata.conf --name fregata xuqingfeng/fregata
```