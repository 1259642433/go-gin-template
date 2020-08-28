FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/project/go-gin-template
COPY . $GOPATH/src/project/go-gin-template
RUN go build .

EXPOSE 12577
ENTRYPOINT ["./go-gin-template"]