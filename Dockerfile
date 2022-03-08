FROM golang:1.17

COPY . /

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build /cmd/http/main.go
EXPOSE 8080

CMD ["./main"]