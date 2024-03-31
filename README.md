# fib_api

## Overview

株式会社Speee様 2024年度夏インターン 参加課題

コードの概要は[こちら](https://github.com/GoRuGoo/fib_api/blob/master/overview.md)

成果物の確認できるURLはこちら

https://fibonacci.gorugoo.com

https://fib-cloud-tokyo-laey36n37q-an.a.run.app

## Requirement

### OS

- Mac OS Sonoma 14.4(動作確認済み)

### Library

- Go
  - Gin
- Docker
- docker-compose

## Installation(local)

1. Clone this repository

```
git clone git@github.com:git@github.com:GoRuGoo/fib_api.git
```
2. Build

```
docker compose up -d
```
3. Adding a self-signed certificate
```
cd build
```
```
mkdir key
```
```
brew install mkcert
```
```
mkcert -install
```
```
mkcert fib-api.com
```


## Usage(local)

1. Build & start container

```
docker compose up -d
```

2. Boot

```
docker compose exec api go run main.go
```
3. Access
```
curl -X GET -H "Content-Type application/json" "https://fib-api.com/fib?n=10"
```




## Author

- [Yuta Ito](https://github.com/GoRuGoo)
