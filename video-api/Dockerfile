FROM golang:1.20.3-buster as dev

WORKDIR /home/app

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ffmpeg 

WORKDIR /home/app
COPY . .

ENV HTTP_PORT=8080 \
    LOG_LEVEL=debug

RUN go install github.com/cosmtrek/air@v1.29.0 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2 && \
    go install github.com/swaggo/swag/cmd/swag@v1.8.12 && \
    go install github.com/go-delve/delve/cmd/dlv@v1.20.1 && \
    go mod download

CMD ["air", "-c", ".air.toml"]



