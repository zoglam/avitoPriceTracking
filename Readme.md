# Avito price tracking

## Wild imports
```bash
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/mux
```

## Run
Parse protection skips with golang 1.10.4 (net/http package on 1.10.4 allows to parse without banning ip)
```bash
# Default
go run main.go
```

## Docker [<a href="https://hub.docker.com/r/zoglam/pricetracking">Docker hub</a>]
```bash
# Build for 1.10.4 golang
sudo docker build -t pricetracking .

# Run
sudo docker run -it -p 8080:8081 pricetracking
```

## Structure
### DB
[<img src="https://live.staticflickr.com/65535/50482812081_682806a9ef_c.jpg" width=405>](https://live.staticflickr.com/65535/50482812081_682806a9ef_c.jpg)
### Programm
[<img src="https://live.staticflickr.com/65535/50490499813_2a8b8044f3_k.jpg" width=405>](https://live.staticflickr.com/65535/50490499813_2a8b8044f3_k.jpg)
