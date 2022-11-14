FROM golang:latest

RUN mkdir /server

WORKDIR /server

COPY  ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN go build ./cmd/http/main.go

CMD ["./main"]


