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

CMD ["/bin/app", "./ports.json"]