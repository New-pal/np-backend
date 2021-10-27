FROM golang:1.17

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod download
COPY . .

RUN apt update
RUN apt install -y build-essential fish

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build='go build main.go' --command=./main
