FROM golang:1.23.1-alpine

WORKDIR /app

COPY . /app

RUN go mod tidy

RUN go install github.com/githubnemo/CompileDaemon@latest

CMD ["CompileDaemon", "--build=go build -o build-bin/main ./src", "--command=./build-bin/main"]
