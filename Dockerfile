FROM golang:1.10.4

LABEL maintainer="<https://github.com/zoglam>"

WORKDIR /app

RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/mattn/go-sqlite3

COPY . .

ENV IP 95.165.148.222
ENV PORT 8081

RUN go build

CMD ["./app"]