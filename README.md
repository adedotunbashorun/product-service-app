## Description
git clone https://github.com/adedotunbashorun/product-service-app.git
cd product-service-app

## Installation

```bash
$ go mod tidy
```

## Running the app

```bash
# development & production mode
$ go run ./cmd/server/main.go

```bash
# development & production mode with docker
$ docker-compose up --build

```bash
# to update swagger documentation
$ swag init -g cmd/server/main.go