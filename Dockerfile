FROM golang:1.20-alpine AS builder

RUN go env -w GO111MODULE=auto \
 && go env -w CGO_ENABLED=0 \
 && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY ./ .

RUN cd /build \
 && GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /build/main /app/

COPY --from=builder /build/env /app/env

WORKDIR /app

RUN chmod +x main

EXPOSE 8088

ENTRYPOINT ["./main"]
