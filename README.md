# fregata
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/xuqingfeng/fregata)
[![api](https://img.shields.io/badge/docs-API-orange.svg?style=flat-square)](https://xuqingfeng.github.io/fregata/api.html)

### Install

```bash
go get github.com/xuqingfeng/fregata/cmd/...
```

### Supported Services

[examples](./examples)

- slack
- telegram
- twilio
- wechat
- smtp
- macos

### Usage

- [fregata.conf](./etc/fregata.conf)
- `fregata -config fregata.conf`
- invoke [API](https://xuqingfeng.github.io/fregata/api.html)

### Test

```bash
make test
```

### Docker

```bash
docker run --rm -p 2017:2017 -v `pwd`/fregata.conf:/etc/fregata/fregata.conf --name fregata xuqingfeng/fregata
```