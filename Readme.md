# Avito price tracking

## wild imports
```bash
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/mux
```

## Run
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
sudo docker run -p 8080:8080 pricetracking
```