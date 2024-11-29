FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN if [ -f config.yml.example ]; then cp config.yml.example config.yml; fi

RUN export GOPROXY=https://goproxy.cn,direct && go mod tidy && go build -ldflags="-s -w" -o ./bin/server ./cmd/server/main.go

EXPOSE 8000

CMD ["./bin/server"]
