FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy && go build -ldflags="-s -w" -o ./bin/server ./cmd/server/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/bin/server .
COPY --from=builder /app/config.yml.example ./config.yml

EXPOSE 8000

CMD ["./server"]
