FROM golang:1.13 as builder

WORKDIR /app

COPY ./webserver.go .
COPY ./build build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./server webserver.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/server .
COPY --from=builder app/build build

ENTRYPOINT ["./server"]
