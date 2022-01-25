FROM golang:latest

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/admin
COPY . $GOPATH/src/admin

RUN go build .

EXPOSE 7070

ENTRYPOINT ["./VideoWeb"]
