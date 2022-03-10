FROM golang:1.17

COPY . /app

WORKDIR /

COPY go.mod /app
COPY go.sum /app

RUN cd /app && go mod download

COPY . .

RUN cd /app && make server

EXPOSE 8080

