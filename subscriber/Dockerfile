FROM golang

WORKDIR /go/bin

RUN go get github.com/canthefason/go-watcher

RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/app

COPY . .

RUN go get -d -v .

RUN go install -v .

EXPOSE 8080

CMD ["watcher"]