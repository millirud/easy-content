FROM golang:1.20.2-buster AS dev

RUN  wget -q -nv -O /bin/hadolint https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
    chmod 777 /bin/hadolint

WORKDIR /home/app
COPY . .

ENV HTTP_PORT=8080 \
    LOG_LEVEL=debug

RUN go install github.com/cosmtrek/air@v1.29.0 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2 && \
    go install github.com/swaggo/swag/cmd/swag@v1.8.10 && \
    go install github.com/go-delve/delve/cmd/dlv@v1.20.1 && \
    go mod download

CMD ["air", "-c", ".air.toml"]
