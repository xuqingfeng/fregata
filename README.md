# fregata
[![travis](https://img.shields.io/travis/xuqingfeng/fregata/master.svg?style=flat-square)](https://travis-ci.org/xuqingfeng/fregata)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/xuqingfeng/fregata)
[![api](https://img.shields.io/badge/docs-API-orange.svg?style=flat-square)](https://xuqingfeng.github.io/fregata/api.html)

### Install

```bash
go get github.com/xuqingfeng/fregata/cmd/fregata
go get github.com/xuqingfeng/fregata/cmd/fregatad
```

### Supported Services

- macos
- slack
- smtp
- telegram
- wechat

### Usage

- prepare a config file
- start `fregata daemon`
- send messages to the API endpoints

### Test

```bash
make test
```

### Docker

```bash
docker run -d -p 2017:2017 -v fregata.conf:/etc/fregata/fregata.conf --name fregata xuqingfeng/fregata 
```