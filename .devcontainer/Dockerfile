FROM golang:1.19

WORKDIR /workspace

RUN apt-get update && \
    apt-get install -y \
    build-essential \
    git \
    curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

COPY . .

EXPOSE 8000

CMD ["/bin/bash"]
