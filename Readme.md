# Avito price tracking

## Wild imports
```bash
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/mux
```

## Run
Parse protection skips on ubuntu 18.04.5 with golang 1.10.4 (net/http package on 1.10.4 allows to parse without banning ip)
```bash
# Default
go run main.go [params]

# Run with params
go run main.go --reset
```

## Docker
```bash
# Build
sudo docker build -t pricetracking .

# Run
sudo docker run --network="host" pricetracking
```