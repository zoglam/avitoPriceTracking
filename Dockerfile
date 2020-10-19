FROM golang:1.10.4

LABEL maintainer="<https://github.com/zoglam>"

WORKDIR /app

RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/mattn/go-sqlite3

COPY . .

ENV IP localhost
ENV PORT 8080

RUN go build

CMD ["./app"]