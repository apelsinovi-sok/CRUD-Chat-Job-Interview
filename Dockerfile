FROM golang:1.17

COPY . /app

WORKDIR /

COPY go.mod /app
COPY go.sum /app

RUN cd /app && go mod download

COPY . .

RUN go build /cmd/http/main.go

EXPOSE 8080

CMD ["./main"]