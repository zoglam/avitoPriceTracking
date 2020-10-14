FROM golang:latest

LABEL maintainer="<https://github.com/zoglam>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV IP localhost
ENV PORT 8080

RUN go build

CMD ["./avitopricetracking"]