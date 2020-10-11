FROM golang:latest

LABEL maintainer="Quique <https://github.com/zoglam>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD ["./avitopricetracking"]