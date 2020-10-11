# Avito price tracking

## Wild imports
```bash
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/mux
```

## Run
Parse protection skips on ubuntu 18.04.5
```bash
# Default
go run main.go [params]

# Run with params
go run main.go --reset
```

## Docker
```bash
# Build
sudo docker build -f DockerFile -t pricetracking .

# Run
sudo docker run --network="host" pricetracking
```