FROM golang:latest

RUN mkdir -p /go/src/github.com/alallema/picture_dictionnary.git/api
WORKDIR /go/src/github.com/alallema/picture_dictionnary.git/api/

ENV GOPATH=/go
ENV GO111MODULE=on

ENV PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:/usr/lib/pkgconfig/

ADD . .

RUN go build -o main cmd/main.go

CMD ["./main"]