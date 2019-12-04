FROM golang:1.13 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/server ./server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build /app

ENV PORT 9090

ENTRYPOINT ["./server"]