FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY / /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /football-subgraph

EXPOSE 8080

ENV GIN_MODE=release
CMD [ "/football-subgraph" ]
