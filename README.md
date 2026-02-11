# fregata
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/xuqingfeng/fregata)

### Install

```bash
go get github.com/xuqingfeng/fregata/cmd/...
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