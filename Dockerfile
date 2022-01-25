FROM golang:latest

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/VideoWeb
COPY . $GOPATH/src/VideoWeb

RUN go build .

EXPOSE 9090

ENTRYPOINT ["./VideoWeb"]
