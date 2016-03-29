FROM golang:1.6.0

RUN mkdir -p /go/src/github.com/hajhatten/graphite-beacon-web
ADD . /go/src/github.com/hajhatten/graphite-beacon-web
WORKDIR /go/src/github.com/hajhatten/graphite-beacon-web

ENV HOME /go/src/github.com/hajhatten/graphite-beacon-web
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 3000

CMD gin
