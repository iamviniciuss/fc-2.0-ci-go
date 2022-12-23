FROM golang:1.18-alpine AS builder

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN apk add git

WORKDIR /build

COPY go.mod ./
RUN go mod tidy
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o app ./src

FROM scratch

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder ["/build/app", "/"]

ENTRYPOINT ["/app"]