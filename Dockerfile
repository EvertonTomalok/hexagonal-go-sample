# Builder
FROM golang:1.23.6-bookworm AS builder
WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download -x


# Testing
FROM builder as testing
COPY . ./
CMD ["go", "test", "./..."]


# Compiler
FROM builder as compiler
COPY . ./

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s" -o /bin/app cmd/application/main.go


# RELEASE
FROM alpine:3.21 AS release

COPY --from=compiler /bin/app /bin/app
COPY ./resources/ports.json ./

COPY .local_env /etc/.local_env
RUN set -o allexport; source /etc/.local_env; set +o allexport

CMD ["/bin/app", "./ports.json"]