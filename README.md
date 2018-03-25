# fregata
[![travis](https://img.shields.io/travis/xuqingfeng/fregata/master.svg?style=flat-square)](https://travis-ci.org/xuqingfeng/fregata)
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
- wechat
- smtp
- macos

### Usage

- prepare a [fregata.conf](./etc/fregata.conf)
- run `fregatad -config fregata.conf`
- send messages to [API endpoints](https://xuqingfeng.github.io/fregata/api.html)

### Test

```bash
make test
```

### Docker

```bash
docker run --rm -p 2017:2017 -v `pwd`/fregata.conf:/etc/fregata/fregata.conf --name fregata xuqingfeng/fregata
```